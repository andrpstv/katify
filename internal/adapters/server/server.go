package server

import (
	"context"
	"katify/internal/adapters/app/user"
	postgres "katify/internal/adapters/db"
	"katify/internal/adapters/http"
	authUseCase "katify/internal/application/app/AuthUseCase"
	"katify/internal/config"
	"katify/internal/delivery"
	"katify/pkg/logger"
	sqlc "katify/sqlc/repository/users"
	httpServer "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pressly/goose/v3"
)

type Server struct {
	log        logger.Logger
	cfg        config.Config
	httpServer *httpServer.Server
}

func NewServer(cfg config.Config, log logger.Logger) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

func (s *Server) Run() error {
	defer func() {
		recover()
	}()

	db, err := postgres.NewPostgresClient(s.cfg.Postgres, s.log)
	if err != nil {
		s.log.Errorf("ошибка поднятия БД: %v", err)
		return err
	}
	if db == nil {
		s.log.Errorf("ошибка поднятия БД: получен nil DB", err)
		return err
	}
	if err := db.Ping(); err != nil {
		s.log.Errorf("ошибка проверки соединения с БД: %v", err)
		return err
	}

	// Test query to ensure DB is fully usable
	if _, err := db.Query("SELECT 1"); err != nil {
		s.log.Errorf("ошибка выполнения тестового запроса к БД: %v", err)
		return err
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		s.log.Errorf("Failed to apply migrations: %v", err)
		return err
	}

	queries := sqlc.New(db)
	userRepo := user.NewUserRepositoryImpl(*queries)

	userService := user.NewUserServiceImpl()
	txManager := postgres.NewTxManager(db, queries)
	authUC := authUseCase.NewAuthUseCaseImpl(userRepo, userService, *txManager)

	authHandler := http.NewAuthHandlerImpl(authUC)

	ginServer := delivery.SetupServer(authHandler)

	s.httpServer = &httpServer.Server{
		Addr:    s.cfg.Server.HTTPPort,
		Handler: ginServer,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != httpServer.ErrServerClosed {
			s.log.Errorf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.log.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.log.Errorf("Server Shutdown error: %v", err)
		return err
	}

	s.log.Infof("Server exiting")
	return nil
}
