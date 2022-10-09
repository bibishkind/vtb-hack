package service

import (
	"context"
)

func (s *service) UpdateScore(ctx context.Context, userId int, score int) error {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	score2, err := s.postgres.GetScore(ctx, tx, userId)
	if err != nil {
		return err
	}

	if err = s.postgres.UpdateScore(ctx, tx, userId, score+score2); err != nil {
		return err
	}

	tx.Commit(ctx)
	return nil
}
