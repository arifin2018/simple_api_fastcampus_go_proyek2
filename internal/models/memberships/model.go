package memberships

import "database/sql"

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
type UserModel struct {
	ID         int          `db:"id"`
	Email      *string      `db:"email"`
Password   string       `db:"password"`
	Username   *string      `db:"username"`
	Created_at sql.NullTime `db:"created_at"`
	Updated_at sql.NullTime `db:"updated_at"`
	Created_by *string      `db:"created_by"`
	Updated_by *string      `db:"updated_by"`
}
