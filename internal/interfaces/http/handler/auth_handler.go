package handler

import (
	"errors"
	"net/http"

	"github.com/alves/age-of-genesis/internal/application/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginService *auth.LoginService
}

func NewAuthHandler(loginService *auth.LoginService) *AuthHandler {
	return &AuthHandler{loginService: loginService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req auth.LoginInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.IP = c.ClientIP()

	out, err := h.loginService.Execute(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		case errors.Is(err, auth.ErrBannedUser):
			c.JSON(http.StatusForbidden, gin.H{"error": "account banned"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, out)
}
