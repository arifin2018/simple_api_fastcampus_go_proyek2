package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/configs"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
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
