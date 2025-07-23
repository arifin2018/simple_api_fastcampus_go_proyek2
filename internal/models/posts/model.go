package posts

import "database/sql"

type CreatePostRequest struct {
	PostTitle   string   `json:"postTitle"`
	PostContent string   `json:"postContent"`
	PostHastags []string `json:"postHastags"`
}

type PostModel struct {
	ID          int64        `db:"id"`
	UserID      int64        `db:"userID"`
	PostTitle   string       `db:"post_title"`
	PostContent string       `db:"post_content"`
	PostHastags string       `db:"post_hastags"`
	Created_at  sql.NullTime `db:"created_at"`
	Updated_at  sql.NullTime `db:"updated_at"`
	Created_by  *string      `db:"created_by"`
	Updated_by  *string      `db:"updated_by"`
}

type GetAllPostDataResponse struct {
	Data       []Post     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Post struct {
	ID          int64    `json:"id"`
	UserId      int64    `json:"userId"`
	Username    string   `json:"username"`
	PostTitle   string   `json:"postTitle"`
	PostContent string   `json:"postContent"`
	PostHastags []string `json:"postHastags"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
