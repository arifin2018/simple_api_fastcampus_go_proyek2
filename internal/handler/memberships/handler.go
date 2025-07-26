package memberships

import (
	"proyek3-catalog-music/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

type service interface {
	SignUp(request memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, s service) *Handler {
	return &Handler{
		Engine:  api,
		service: s,
	}
}

func (h *Handler) RegisterRoutes() {
	membership := h.Group("/memberships")
	membership.POST("/memberships/signup", h.SignUp)
}
