package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Transactions
	Users
	RefreshTokens
}

type Transactions interface {
	AcquireTx(ctx context.Context) (pgx.Tx, error)
}

type Users interface {
	CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) error
	GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*domain.User, error)
}

type RefreshTokens interface {
	CreateRefreshToken(ctx context.Context, tx pgx.Tx, token string) error
}

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) *Postgres {
	return &Postgres{
		pool: pool,
	}
}
