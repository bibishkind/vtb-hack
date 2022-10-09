package service

import (
	vtb2 "coffee-layered-architecture/api/vtb"
	"coffee-layered-architecture/internal/domain"
	postgres2 "coffee-layered-architecture/internal/postgres"
	"coffee-layered-architecture/pkg/auth"
	"context"
)

type Service interface {
	Auth
	Finance
	Cards
	Profile
	Scores
	Tasks
}

type Auth interface {
	SignUp(ctx context.Context, user *domain.User) error
	SignIn(ctx context.Context, user *domain.User) (string, error)
	ParseAccessToken(accessToken string) (int, error)
}

type Finance interface {
	GetBalance(ctx context.Context, userId int) (float32, float32, error)
	TransferMatic(ctx context.Context, senderId int, receiverId int, amount float32) (string, error)
	TransferRuble(ctx context.Context, senderId int, receiverId int, amount float32) (string, error)
}

type Cards interface {
	CreateCard(ctx context.Context, userId int, card *domain.Card) (int, error)
	GetAllCards(ctx context.Context) ([]*domain.Card, error)
	DeleteCard(ctx context.Context, userId int, cardId int) error
}

type Tasks interface {
	CreateTask(ctx context.Context, userId int, task *domain.Task) (int, error)
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
}

type Profile interface {
	GetProfile(ctx context.Context, userId int) (*domain.Profile, error)
}

type Scores interface {
	UpdateScore(ctx context.Context, userId int, score int) error
}

type service struct {
	postgres     postgres2.Postgres
	tokenManager *auth.TokenManager
	vtb          vtb2.Vtb
}

func NewService(postgres postgres2.Postgres, tokenManager *auth.TokenManager, vtb vtb2.Vtb) Service {
	return &service{
		postgres:     postgres,
		tokenManager: tokenManager,
		vtb:          vtb,
	}
}
