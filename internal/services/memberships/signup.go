package memberships

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepository.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("username or email already exist")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// return errors.New("username or email already exist")
	now := time.Now()
	model := memberships.UserModel{
		Email:      &req.Email,
		Password:   string(pass),
		Created_at: sql.NullTime{Time: now, Valid: true},
		Updated_at: sql.NullTime{Time: now, Valid: true},
		Created_by: &req.Username,
		Updated_by: &req.Username,
		Username:   &req.Username,
	}
	if err := s.membershipRepository.CreateUser(ctx, &model); err != nil {
		return err
	}
	return nil
}
