package service

import (
	vtb2 "coffee-layered-architecture/api/vtb"
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
	SignIn(ctx context.Context, user *domain.User) (string, error)
	IdentifyUser(ctx context.Context, accessToken string) (int, error)
}

type Service struct {
	postgres     postgres2.Client
	tokenManager *auth.TokenManager
	vtb          vtb2.Vtb
}

func NewService(postgres postgres2.Client, tokenManager *auth.TokenManager, vtb vtb2.Vtb) *Service {
	return &Service{
		postgres:     postgres,
		tokenManager: tokenManager,
		vtb:          vtb,
	}
}
