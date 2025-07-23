package posts

import (
	"context"
	"database/sql"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	query := `SELECT id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by 
	          FROM user_activities 
	          WHERE post_id = ? AND user_id = ?`

	var activity posts.UserActivityModel
	err := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID).Scan(
		&activity.ID,
		&activity.PostID,
		&activity.UserID,
		&activity.IsLiked,
		&activity.CreatedAt,
		&activity.UpdatedAt,
		&activity.CreatedBy,
		&activity.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No activity found
		}
		return nil, err // Other error
	}
	return &activity, nil
}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (post_id, user_id, is_liked, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedBy, model.UpdatedBy)
	return err
}

func (r *repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities SET is_liked = ?, updated_by = ? WHERE post_id = ? AND user_id = ?`

	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedBy, model.PostID, model.UserID)
	return err
}

func (r *repository) CountLikeByPostId(ctx context.Context, postID int64) (int, error) {
	query := `select count(id) from user_activities p where p.id = ? and p.is_liked = true`
	var count int
	if err := r.db.QueryRowContext(ctx, query, postID).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, nil
		}
		return count, err
	}
	return count, nil
}
