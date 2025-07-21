package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id,post_title,post_content,post_hastags,created_at,updated_at,created_by,updated_by) VALUES (?,?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHastags, model.Created_at, model.Updated_at, model.Created_by, model.Updated_by)
	if err != nil {
		return err
	}
	return nil
}
