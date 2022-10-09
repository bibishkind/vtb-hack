package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres interface {
	Transactions
	Users
}

type Transactions interface {
	AcquireTx(ctx context.Context) (pgx.Tx, error)
	RollbackTx(ctx context.Context, tx pgx.Tx) error
	CommitTx(ctx context.Context, tx pgx.Tx) error
}

type Users interface {
	CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) error
	GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*domain.User, error)
	GetUserById(ctx context.Context, tx pgx.Tx, id int) (*domain.User, error)
}

type postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) Postgres {
	return &postgres{
		pool: pool,
	}
}
