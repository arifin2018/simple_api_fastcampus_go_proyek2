package memberships

import (
	"context"
	"database/sql"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `select id,email, password,created_at,updated_at,created_by,updated_by,username from users where email = $1 or username = $2`
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, response.Password, response.Email, response.Created_at, response.Created_by, response.Updated_at, response.Updated_by, response.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, user *memberships.UserModel) error {
	query := `INSERT INTO users(email,password,username,created_at,updated_at,created_by,updated_by) values ($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Username, user.Created_at, user.Updated_at, user.Created_by, user.Updated_by)
	if err != nil {
		return err
	}
	return nil
}
