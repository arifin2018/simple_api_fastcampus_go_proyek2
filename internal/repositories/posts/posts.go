package posts

import (
	"context"
	"strings"

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

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostDataResponse, error) {
	query := `SELECT p.id,p.user_id,u.username,p.post_title,p.post_content,p.post_hastags,p.created_at,p.updated_at,p.created_by,p.updated_by FROM posts p LEFT JOIN users u ON p.user_id = u.id order by p.updated_at DESC LIMIT ? OFFSET ? `
	response := posts.GetAllPostDataResponse{}
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return posts.GetAllPostDataResponse{}, err
	}
	defer rows.Close()

	data := []posts.Post{}
	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)
		if err := rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHastags, &model.Created_at, &model.Updated_at, &model.Created_by, &model.Updated_by); err != nil {
			return posts.GetAllPostDataResponse{}, err
		}
		data = append(data, posts.Post{
			ID:          model.ID,
			UserId:      model.UserID,
			Username:    username,
			PostTitle:   model.PostTitle,
			PostContent: model.PostContent,
			PostHastags: strings.Split(model.PostHastags, ","),
		})
	}
	if err := rows.Err(); err != nil {
		return posts.GetAllPostDataResponse{}, err
	}

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	return response, nil
}

func (r *repository) GetPostById(ctx context.Context, postId int64) (posts.Post, error) {
	query := `SELECT p.id,p.user_id,u.username,p.post_title,p.post_content,p.post_hastags,p.created_at,p.updated_at,p.created_by,p.updated_by FROM posts p LEFT JOIN users u ON p.user_id = u.id WHERE p.id = ?`
	var model posts.PostModel
	var username string
	var isLiked bool

	err := r.db.QueryRowContext(ctx, query, postId).Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHastags, &model.Created_at, &model.Updated_at, &model.Created_by, &model.Updated_by)
	if err != nil {
		return posts.Post{}, err
	}

	response := posts.Post{
		ID:          model.ID,
		UserId:      model.UserID,
		Username:    username,
		PostTitle:   model.PostTitle,
		PostContent: model.PostContent,
		PostHastags: strings.Split(model.PostHastags, ","),
		IsLiked:     isLiked,
	}

	return response, nil
}
