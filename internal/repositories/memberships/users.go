package memberships

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string, userID int) (*memberships.UserModel, error) {
	query := `select id,email, password,created_at,updated_at,created_by,updated_by,username from users where email = ? or username = ? or id = ?`
	row := r.db.QueryRowContext(ctx, query, email, username, userID)

	var response memberships.UserModel
	err := row.Scan(
		&response.ID,
		&response.Email,
		&response.Password,
		&response.Created_at,
		&response.Updated_at,
		&response.Created_by,
		&response.Updated_by,
		&response.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, user *memberships.UserModel) error {
	query := `INSERT INTO users(email,password,username,created_at,updated_at,created_by,updated_by) values (?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Username, user.Created_at, user.Updated_at, user.Created_by, user.Updated_by)
	if err != nil {
		fmt.Println("2")
		return err
	}
	return nil
}
