package memperships

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/middleware"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userId int, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipService membershipService
}

func NewHandler(api *gin.Engine, membershipService membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	route.POST("/refresh", middleware.AuthRefreshMiddleware(), h.RefreshToken)
}
