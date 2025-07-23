package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (r *service) GetPostById(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postByID, err := r.postRepo.GetPostById(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get post by id")
		return nil, err
	}
	countLike, err := r.postRepo.CountLikeByPostId(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("failed to count likes by post id")
		return nil, err
	}
	comments, err := r.postRepo.GetCommentByPostID(ctx, int(postID))
	if err != nil {
		log.Error().Err(err).Msg("failed to get comments by post id")
		return nil, err
	}

	response := posts.GetPostResponse{
		PostDetail: postByID,
		LikedCount: countLike,
		Comments:   comments,
	}
	return &response, nil
}
