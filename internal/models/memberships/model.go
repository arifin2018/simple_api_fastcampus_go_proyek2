package memberships

import "time"

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	ID         int        `db:"id"`
	Email      string     `db:"email"`
	Password   string     `db:"password"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
	Created_by string     `db:"created_by"`
	Updated_by string     `db:"updated_by"`
	Username   string     `db:"username"`
}
