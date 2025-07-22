package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/middleware"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postId, userId int, request posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine
	postService postService
}

func NewHandler(api *gin.Engine, postService postService) *Handler {
	return &Handler{
		Engine:      api,
		postService: postService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())
	route.POST("create", h.CreatePost)
	route.POST("comment/:postID", h.CreateComment)
	route.PUT("user_activity/:postID", h.UpsertUserActivity)
}
