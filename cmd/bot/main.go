package main

import (
	"database/sql"
	"log"
	"os"
	"report/internal/adapters/app/user"
	authUseCase "report/internal/application/app/AuthUseCase"
	sqlc "report/sqlc/repository/users"

	"report/internal/adapters/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/report_parser?sslmode=disable"
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	userRepo := user.NewUserRepositoryImpl(*queries)
	userService := user.NewUserServiceImpl()
	authUC := authUseCase.NewAuthUseCase(userRepo, userService)
	authHandler := http.NewAuthHandler(authUC)

	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
