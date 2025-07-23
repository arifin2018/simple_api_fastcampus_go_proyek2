package tokenGenerate

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
