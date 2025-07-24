package memberships

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/jwt"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/tokenGenerate"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepository.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("user not found")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", errors.New("email or password is invalid")
	}

	token, err := jwt.CreateToken(int64(user.ID), *user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		return "", "", err
	}

	existingToken, err := s.membershipRepository.GetRefreshToken(ctx, int64(user.ID), time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", "", err
	}

	if existingToken != nil && existingToken.RefreshToken != "" {
		log.Info().Msg("refresh token already exists, skipping creation")
		return token, existingToken.RefreshToken, nil
	}

	refresh_token := tokenGenerate.GenerateRefreshToken()
	if refresh_token == "" {
		log.Error().Msg("failed to generate refresh token")
		return token, "", errors.New("failed to generate refresh token")
	}
	fmt.Printf("refresh aja\n")
	fmt.Printf("%+v\n", time.Now())
	err = s.membershipRepository.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refresh_token,
		ExpiredAt:    sql.NullTime{Time: time.Now().Add(10 * time.Minute), Valid: true}, // 10 menit
		CreatedBy:    user.Username,
		UpdatedBy:    user.Username,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to insert refresh token")
		return token, refresh_token, err
	}

	return token, refresh_token, nil
}
