package server

import (
	"context"
	"log"
	httpServer "net/http"
	"os"
	"os/signal"
	"report/internal/adapters/app/user"
	postgres "report/internal/adapters/db"
	"report/internal/adapters/http"
	authUseCase "report/internal/application/app/AuthUseCase"
	"report/internal/config"
	"report/internal/delivery"
	"report/pkg/logger"
	sqlc "report/sqlc/repository/users"
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
	db, _ := postgres.NewPostgresClient(s.cfg.Postgres)
	queries := sqlc.New(db)
	userRepo := user.NewUserRepositoryImpl(*queries)

	userService := user.NewUserServiceImpl()
	txManager := postgres.NewTxManager(db, queries)
	authUC := authUseCase.NewAuthUseCaseImpl(userRepo, userService, *txManager)

	authHandler := http.NewAuthHandlerImpl(authUC)

	ginServer := delivery.SetupServer(authHandler)

	if err := goose.Up(db, "./migrations"); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

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
		return err
	}

	s.log.Infof("Server exiting")
	return nil
}
