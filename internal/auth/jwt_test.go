package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	secret := "secret"
	os.Setenv("JWT_SECRET_TOKEN", secret)
	id, email := "1", "kaungminhtet@example.com"
	tokenString, err := generateToken(id, email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	assert.NotNil(t, err)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Errorf("Type assertion error")
	}

	assert.Equal(t, id, claims["jti"])
	assert.Equal(t, email, claims["email"])
}
