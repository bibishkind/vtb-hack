package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) CreateScore(ctx context.Context, tx pgx.Tx, userId int) error {
	q1 := `insert into scores (user_id) values ($1)`
	_, err := tx.Exec(ctx, q1, userId)
	return err
}

func (p *postgres) UpdateScore(ctx context.Context, tx pgx.Tx, userId int, score int) error {
	q1 := `update scores set score=$1 where user_id=$2`
	_, err := tx.Exec(ctx, q1, score, userId)
	return err
}

func (p *postgres) GetScore(ctx context.Context, tx pgx.Tx, userId int) (int, error) {
	q1 := `select score from scores where user_id=$1`
	row := tx.QueryRow(ctx, q1, userId)
	var score int
	if err := row.Scan(&score); err != nil {
		return 0, err
	}
	return score, nil
}
