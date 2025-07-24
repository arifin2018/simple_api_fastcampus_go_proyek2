package memberships

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/models/memberships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userId int, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepository.GetRefreshToken(ctx, int64(userId), time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", err
	}
	if existingRefreshToken == nil || existingRefreshToken.RefreshToken != request.Token {
		log.Error().Msg("refresh token not found or does not match")
		return "", errors.New("invalid refresh token")
	}

	user, err := s.membershipRepository.GetUser(ctx, "", "", userId)
	if err != nil {
		fmt.Println("123")
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		log.Error().Msg("user not found")
		return "", errors.New("user not found")
	}

	if existingRefreshToken.ExpiredAt.Time.Before(time.Now()) {
		log.Error().Msg("refresh token has expired")
		return "", errors.New("refresh token has expired")
	}
	newAccessToken, err := jwt.CreateToken(int64(userId), *existingRefreshToken.CreatedBy, s.cfg.Service.SecretJwt)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}
