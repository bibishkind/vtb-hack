package auth

import (
	"time"
)

type TokenManager struct {
	SigningKey string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func NewTokenManager(signingKey string, accessTTL, refreshTTL time.Duration) *TokenManager {
	return &TokenManager{
		SigningKey: signingKey,
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}
}
