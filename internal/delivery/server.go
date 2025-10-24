package delivery

import (
	"report/internal/adapters/http"
	authUseCase "report/internal/application/app/AuthUseCase"

	"github.com/gin-gonic/gin"
)

func SetupServer(authUseCase *authUseCase.AuthUseCaseImpl) *gin.Engine {
	router := gin.Default()
	authHandler := http.NewAuthHandler(authUseCase)

	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	return router
}
