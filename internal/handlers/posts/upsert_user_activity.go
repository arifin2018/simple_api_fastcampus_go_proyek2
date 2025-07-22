package posts

import (
	"net/http"
	"strconv"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	// Call the service method to upsert user activity
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postID := c.Param("postID")
	postid, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + " postID",
		})
		return
	}

	userID := c.Query("userID")
	userid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + " userid",
		})
		return
	}

	if err := h.postService.UpsertUserActivity(ctx, postid, userid, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User activity upserted successfully",
		"data":    request,
	})
}
