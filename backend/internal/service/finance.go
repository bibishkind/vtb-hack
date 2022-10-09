package service

import (
	"context"
)

func (s *service) GetBalance(ctx context.Context, userId int) (float32, float32, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	user, err := s.postgres.GetUserById(ctx, tx, userId)
	if err != nil {
		return 0, 0, err
	}

	s.postgres.CommitTx(ctx, tx)
	return s.vtb.GetBalance(user.PublicKey)
}

func (s *service) TransferMatic(ctx context.Context, senderId int, receiverId int, amount float32) (string, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return "", err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	sender, err := s.postgres.GetUserById(ctx, tx, senderId)
	if err != nil {
		return "", err
	}

	receiver, err := s.postgres.GetUserById(ctx, tx, receiverId)
	if err != nil {
		return "", err
	}

	senderPrivateKey := sender.PrivateKey
	receiverPublicKey := receiver.PrivateKey

	hash, err := s.vtb.TransferMatic(senderPrivateKey, receiverPublicKey, amount)
	if err != nil {
		return "", err
	}

	tx.Commit(ctx)
	return hash, nil
}

func (s *service) TransferRuble(ctx context.Context, senderId int, receiverId int, amount float32) (string, error) {
	tx, err := s.postgres.AcquireTx(ctx)
	if err != nil {
		return "", err
	}
	defer s.postgres.RollbackTx(ctx, tx)

	sender, err := s.postgres.GetUserById(ctx, tx, senderId)
	if err != nil {
		return "", err
	}

	receiver, err := s.postgres.GetUserById(ctx, tx, receiverId)
	if err != nil {
		return "", err
	}

	senderPrivateKey := sender.PrivateKey
	receiverPublicKey := receiver.PrivateKey

	hash, err := s.vtb.TransferRuble(senderPrivateKey, receiverPublicKey, amount)
	if err != nil {
		return "", err
	}

	tx.Commit(ctx)
	return hash, nil
}
