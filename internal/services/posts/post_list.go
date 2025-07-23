package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostDataResponse, error) {
	limit := pageSize
	offset := pageIndex * (pageSize - 1)
	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to get all posts")
		return posts.GetAllPostDataResponse{}, err
	}

	return response, nil
}
