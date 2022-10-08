package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"time"
)

type RefreshClaims struct {
	Id string
	*jwt.RegisteredClaims
}

func (m *TokenManager) GenerateRefreshToken() (string, error) {
	signingKey := []byte(os.Getenv("SIGNING_KEY"))

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.RefreshTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshClaims{
		Id:               uuid.Must(uuid.NewRandom()).String(),
		RegisteredClaims: claims,
	})

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (m *TokenManager) ParseRefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims.Id, nil
	} else {
		return "", errors.New("failed to parse refresh.go token")
	}
}
