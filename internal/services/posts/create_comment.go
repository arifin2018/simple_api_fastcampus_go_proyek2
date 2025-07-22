package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	userid := strconv.Itoa(int(userId))
	model := posts.CommentModel{
		PostID:         postId,
		UserID:         userId,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      userid,
		UpdatedBy:      userid,
	}
	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
		return err
	}

	return nil
}
