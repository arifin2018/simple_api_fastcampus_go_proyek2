package memperships

import (
	"net/http"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipService.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := memberships.LoginResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, response)
}
