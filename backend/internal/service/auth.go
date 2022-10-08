package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
)

func (s *Service) SignUp(ctx context.Context, user *domain.User) error {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	publicKey, privateKey, err := s.vtb.NewWallet()
	if err != nil {
		return err
	}

	user.PublicKey = publicKey
	user.PrivateKey = privateKey

	if err = s.postgres.CreateUser(ctx, tx, user); err != nil {
		return err
	}

	tx.Commit(ctx)
	return nil
}
func (s *Service) SignIn(ctx context.Context, user *domain.User) (string, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	user, err = s.postgres.GetUserByUsername(ctx, tx, user.Username)
	if err != nil {
		return "", err
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(user.Id)
	if err != nil {
		return "", err
	}

	tx.Commit(ctx)
	return accessToken, nil
}

func (s *Service) IdentifyUser(ctx context.Context, accessToken string) (int, error) {
	return s.tokenManager.ParseAccessToken(accessToken)
}
