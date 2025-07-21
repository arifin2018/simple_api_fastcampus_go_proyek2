package memberships

import (
	"context"
	"errors"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepository.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("email or password is invalid")
	}

	token, err := jwt.CreateToken(int64(user.ID), *user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		return "", err
	}
	return token, nil
}
