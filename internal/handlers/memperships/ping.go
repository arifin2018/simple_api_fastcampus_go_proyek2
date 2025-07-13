package memperships

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ping(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "pong",
	})
}
