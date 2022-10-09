package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
)

func (s *service) CreateCard(ctx context.Context, userId int, card *domain.Card) (int, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return 0, err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	user, err := s.postgres.GetUserById(ctx, tx, userId)
	if err != nil {
		return 0, err
	}

	if user.Role != domain.UserRoleAdmin {
		return 0, errors.New("user must have admin role")
	}

	cardId, err := s.postgres.CreateCard(ctx, tx, card)
	if err != nil {
		return 0, err
	}

	s.postgres.CommitTx(ctx, tx)
	return cardId, nil
}
