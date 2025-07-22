package posts

import "time"

type UserActivityRequest struct {
	IsLiked bool `json:"isLiked"`
}

type UserActivityModel struct {
	ID        int       `db:"id"`
	PostID    int       `db:"post_id"`
	UserID    int       `db:"user_id"`
	IsLiked   bool      `db:"is_liked"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}
