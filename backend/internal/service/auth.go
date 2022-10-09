package service

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
)

func (s *service) SignUp(ctx context.Context, user *domain.User) error {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	publicKey, privateKey, err := s.vtb.CreateWallet()
	if err != nil {
		return err
	}

	user.PublicKey = publicKey
	user.PrivateKey = privateKey

	user.Password = s.hasher.HashSha256(user.Password)

	userId, err := s.postgres.CreateUser(ctx, tx, user)
	if err != nil {
		return err
	}

	if err = s.postgres.CreateScore(ctx, tx, userId); err != nil {
		return err
	}

	s.postgres.CommitTx(ctx, tx)
	return nil
}

func (s *service) SignIn(ctx context.Context, user *domain.User) (string, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return "", err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	user2, err := s.postgres.GetUserByUsername(ctx, tx, user.Username)
	if err != nil {
		return "", err
	}

	if s.hasher.HashSha256(user.Password) != user2.Password {
		return "", errors.New("wrong password")
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(user2.Id)
	if err != nil {
		return "", err
	}

	s.postgres.CommitTx(ctx, tx)
	return accessToken, nil
}

func (s *service) ParseAccessToken(accessToken string) (int, error) {
	return s.tokenManager.ParseAccessToken(accessToken)
}
