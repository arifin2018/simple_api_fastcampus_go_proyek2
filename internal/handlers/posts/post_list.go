package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPosts(c *gin.Context) {
	ctx := c.Request.Context()
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("pageIndex"))
	if err != nil || pageIndex < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page index"})
		return
	}

	response, err := h.postService.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetPostById(c *gin.Context) {
	ctx := c.Request.Context()
	postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
	if err != nil || postID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	response, err := h.postService.GetPostById(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
