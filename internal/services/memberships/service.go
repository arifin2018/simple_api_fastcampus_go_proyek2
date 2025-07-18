package memberships

import (
	"context"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user *memberships.UserModel) error
}

type service struct {
	membershipRepository membershipRepository
}

func NewService(membershipRepository membershipRepository) *service {
	return &service{
		membershipRepository: membershipRepository,
	}
}
