package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/pkg/response"
	"github.com/bo-patrol/internal/repository"
	"github.com/bo-patrol/internal/service"
)

var (
	authService *service.AuthService
)

func initAuth() {
	userRepo := repository.NewUserRepository(db.DB)
	authService = service.NewAuthService(userRepo)
}

func Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	user, err := authService.Register(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, user.ToResponse())
}

func Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := authService.Login(&req)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, resp)
}
