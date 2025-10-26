package http

import (
	"net/http"

	authUseCase "report/internal/application/app/AuthUseCase"
	dto "report/internal/dto/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *authUseCase.AuthUseCaseImpl
}

func NewAuthHandlerImpl(usecase *authUseCase.AuthUseCaseImpl) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.Register(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registration successful",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.usecase.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
