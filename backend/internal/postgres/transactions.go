package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) AcquireTx(ctx context.Context) (pgx.Tx, error) {
	return p.pool.Begin(ctx)
}

func (p *Postgres) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}

func (p *Postgres) CommitTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}
