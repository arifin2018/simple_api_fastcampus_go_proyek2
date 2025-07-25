package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/configs"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostDataResponse, error)
	GetPostById(ctx context.Context, postId int64) (posts.Post, error)
	CountLikeByPostId(ctx context.Context, postID int64) (int, error)
	GetCommentByPostID(ctx context.Context, postId int) ([]posts.Comment, error)
}

type service struct {
	cfg      configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      *cfg,
		postRepo: postRepo,
	}
}
