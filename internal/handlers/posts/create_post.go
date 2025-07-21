package posts

import (
	"net/http"
	"strconv"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()
	var req posts.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	userId := c.Query("userId")
	userid, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	if err := h.postService.CreatePost(ctx, int64(userid), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
