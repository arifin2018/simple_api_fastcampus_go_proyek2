package posts

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments(post_id,user_id,comment_content,created_at,updated_at,created_by,updated_by) VALUES (?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCommentByPostID(ctx context.Context, postId int) ([]posts.Comment, error) {
	query := `SELECT c.id,c.post_id,c.user_id,c.comment_content,u.username FROM comments c JOIN users u ON c.user_id = u.id where c.post_id = ?`
	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []posts.Comment
	for rows.Next() {
		var comment posts.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserId, &comment.CommentContent, &comment.Username); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
