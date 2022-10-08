package mocks

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type Postgres struct {
	mock.Mock
}

func (p *Postgres) AcquireTx(ctx context.Context) (pgx.Tx, error) {
	args := p.Called(mock.Anything)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

func (p *Postgres) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	args := p.Called(mock.Anything, mock.Anything)
	return args.Error(0)
}

func (p *Postgres) CommitTx(ctx context.Context, tx pgx.Tx) error {
	args := p.Called(mock.Anything, mock.Anything)
	return args.Error(0)
}
