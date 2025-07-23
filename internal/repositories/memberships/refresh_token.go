package memberships

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_token (user_id, refresh_token, expired_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedBy, model.UpdatedBy)
	return err
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error) {
	query := `SELECT id,user_id,refresh_token,expired_at,created_at,updated_at,created_by,updated_by FROM refresh_token WHERE user_id = ? and expired_at >= ?`
	var response memberships.RefreshTokenModel
	err := r.db.QueryRowContext(ctx, query, userID, now).Scan(&response.ID, &response.UserID, &response.RefreshToken, &response.ExpiredAt, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return &memberships.RefreshTokenModel{}, nil // No token found
		}
		return &memberships.RefreshTokenModel{}, err // Other error
	}
	return &response, nil
}
