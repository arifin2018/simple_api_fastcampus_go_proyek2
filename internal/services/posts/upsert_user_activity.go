package posts

import (
	"context"
	"errors"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postId, userId int, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postId,
		UserID:    userId,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: userId,
		UpdatedBy: userId,
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user activity")
		return err
	}

	if userActivity == nil {
		// Insert new user activity
		if !request.IsLiked {
			return errors.New("cannot insert user activity with isLiked false")
		}
		if err := s.postRepo.CreateUserActivity(ctx, model); err != nil {
			log.Error().Err(err).Msg("failed to create user activity")
			return err
		}
	} else {
		// Update existing user activity
		if err := s.postRepo.UpdateUserActivity(ctx, model); err != nil {
			log.Error().Err(err).Msg("failed to update user activity")
			return err
		}
	}

	return nil
}
