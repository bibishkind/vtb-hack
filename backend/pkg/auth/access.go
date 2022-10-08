package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AccessClaims struct {
	Id int
	*jwt.RegisteredClaims
}

func (m *TokenManager) GenerateAccessToken(id int) (string, error) {
	signingKey := []byte(m.SigningKey)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessClaims{
		Id:               id,
		RegisteredClaims: claims,
	})

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (m *TokenManager) ParseAccessToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*AccessClaims); ok && token.Valid {
		return claims.Id, nil
	} else {
		return 0, errors.New("failed to parse access token")
	}
}
