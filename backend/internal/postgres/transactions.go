package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) AcquireTx(ctx context.Context) (pgx.Tx, error) {
	return p.pool.Begin(ctx)
}

func (p *postgres) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}

func (p *postgres) CommitTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}
