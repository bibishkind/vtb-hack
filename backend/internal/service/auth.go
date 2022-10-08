package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"log"
)

func (s *Service) SignUp(ctx context.Context, user *domain.User) error {
	tx, err := s.postgres.AcquireTx(ctx)

	if err != nil || tx == nil {
		log.Println(err)
		return err
	}

	if err = s.postgres.CreateUser(ctx, tx, user); err != nil {
		log.Println(err)
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)
	return nil
}
func (s *Service) SignIn(ctx context.Context, user *domain.User) (string, string, error) {
	tx, err := s.postgres.AcquireTx(ctx)

	if err != nil || tx == nil {
		log.Println(err)
		return "", "", err
	}

	user, err = s.postgres.GetUserByUsername(ctx, tx, user.Username)
	if err != nil {
		tx.Rollback(ctx)
		log.Println(err)
		return "", "", err
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(user.Id)
	if err != nil {
		log.Println(err)
		tx.Rollback(ctx)
		return "", "", err
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken()
	if err != nil {
		log.Println(err)
		tx.Rollback(ctx)
		return "", "", err
	}

	if err = s.postgres.CreateRefreshToken(ctx, tx, refreshToken); err != nil {
		log.Println(err)
		tx.Rollback(ctx)
		return "", "", err
	}

	tx.Commit(ctx)
	return accessToken, refreshToken, nil
}
