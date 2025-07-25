package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		"Username":  username,
		"Id":        id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidationToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid || claims["ExpiresAt"].(float64) < float64(time.Now().Unix()) {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["Id"].(float64)), claims["Username"].(string), nil
}

func ValidationTokenWithoutExpired(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["Id"].(float64)), claims["Username"].(string), nil
}
