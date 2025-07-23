package memberships

import (
	"context"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/configs"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user *memberships.UserModel) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

type service struct {
	cfg                  *configs.Config
	membershipRepository membershipRepository
}

func NewService(cfg *configs.Config, membershipRepository membershipRepository) *service {
	return &service{
		cfg:                  cfg,
		membershipRepository: membershipRepository,
	}
}
