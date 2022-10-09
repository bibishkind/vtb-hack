package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
)

func (s *service) CreateTask(ctx context.Context, userId int, task *domain.Task) (int, error) {
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

	taskId, err := s.postgres.CreateTask(ctx, tx, task)
	if err != nil {
		return 0, err
	}

	s.postgres.CommitTx(ctx, tx)
	return taskId, nil
}

func (s *service) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	cards, err := s.postgres.GetAllTasks(ctx, tx)
	if err != nil {
		return nil, err
	}

	s.postgres.CommitTx(ctx, tx)
	return cards, nil
}
