package delivery

import (
	"report/internal/adapters/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer(authHandler *http.AuthHandler) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
	return router
}
