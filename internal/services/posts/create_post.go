package posts

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/helpers"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHastags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:      userId,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		Created_at:  sql.NullTime{Time: now, Valid: true},
		Updated_at:  sql.NullTime{Time: now, Valid: true},
		Created_by:  helpers.PtrString(strconv.Itoa(int(userId))),
		Updated_by:  helpers.PtrString(strconv.Itoa(int(userId))),
	}
	if err := s.postRepo.CreatePost(ctx, model); err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}
