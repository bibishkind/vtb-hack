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
	Cards
	Scores
	Tasks
}

type Transactions interface {
	AcquireTx(ctx context.Context) (pgx.Tx, error)
	RollbackTx(ctx context.Context, tx pgx.Tx) error
	CommitTx(ctx context.Context, tx pgx.Tx) error
}

type Users interface {
	CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) (int, error)
	GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*domain.User, error)
	GetUserById(ctx context.Context, tx pgx.Tx, id int) (*domain.User, error)
}

type Cards interface {
	CreateCard(ctx context.Context, tx pgx.Tx, card *domain.Card) (int, error)
	GetAllCards(ctx context.Context, tx pgx.Tx) ([]*domain.Card, error)
	DeleteCard(ctx context.Context, tx pgx.Tx, cardId int) error
}

type Tasks interface {
	CreateTask(ctx context.Context, tx pgx.Tx, task *domain.Task) (int, error)
	GetAllTasks(ctx context.Context, tx pgx.Tx) ([]*domain.Task, error)
	DeleteTask(ctx context.Context, tx pgx.Tx, taskId int) error
}

type Scores interface {
	CreateScore(ctx context.Context, tx pgx.Tx, userId int) error
	UpdateScore(ctx context.Context, tx pgx.Tx, userId int, score int) error
	GetScore(ctx context.Context, tx pgx.Tx, userId int) (int, error)
}

type postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(pool *pgxpool.Pool) Postgres {
	return &postgres{
		pool: pool,
	}
}
