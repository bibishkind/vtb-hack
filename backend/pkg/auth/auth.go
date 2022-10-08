package auth

import (
	"time"
)

type TokenManager struct {
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func NewTokenManager(accessTTL, refreshTTL time.Duration) *TokenManager {
	return &TokenManager{
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}
}
