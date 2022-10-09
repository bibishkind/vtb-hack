package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
)

func (s *service) GetProfile(ctx context.Context, userId int) (*domain.Profile, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	user, err := s.postgres.GetUserById(ctx, tx, userId)
	if err != nil {
		return nil, err
	}

	var profile domain.Profile

	profile.Username = user.Username
	profile.FirstName = user.FirstName
	profile.MiddleName = user.MiddleName
	profile.LastName = user.LastName
	profile.Email = user.Email
	profile.Role = user.Role

	s.postgres.CommitTx(ctx, tx)
	return &profile, nil
}
