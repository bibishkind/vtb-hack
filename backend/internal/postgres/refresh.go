package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) CreateRefreshToken(ctx context.Context, tx pgx.Tx, token string) error {
	q1 := `insert into refresh_tokens (token) values ($1)`
	_, err := tx.Exec(ctx, q1, token)
	return err
}
