package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/common"
	"os"
	"time"
)

func generateToken(id, email string) (string, error) {
	claims := common.UserClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_TOKEN")
	return token.SignedString([]byte(secret))
}
