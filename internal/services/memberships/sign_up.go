package memberships

import (
	"proyek3-catalog-music/internal/models/memberships"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	user, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err == nil {
		return err
	}

	if user != nil {
		log.Error().Msg("user already exists")
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return err
	}

	model := memberships.User{
		Email:        request.Email,
		Username:     request.Username,
		PasswordHash: string(hashedPassword),
		CreatedBy:    request.Email,
		UpdatedBy:    request.Email,
	}
	return s.repository.CreateUser(&model)
}
