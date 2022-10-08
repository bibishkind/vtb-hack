package service

import (
	"coffee-layered-architecture/internal/domain"
	postgres2 "coffee-layered-architecture/internal/postgres"
	"coffee-layered-architecture/pkg/auth"
	"context"
)

type Client interface {
	Auth
}

type Auth interface {
	SignUp(ctx context.Context, user *domain.User) error
	SignIn(ctx context.Context, user *domain.User) (string, string, error)
}

type Service struct {
	postgres     postgres2.Client
	tokenManager *auth.TokenManager
}

func NewService(postgres postgres2.Client, tokenManager *auth.TokenManager) *Service {
	return &Service{
		postgres:     postgres,
		tokenManager: tokenManager,
	}
}
