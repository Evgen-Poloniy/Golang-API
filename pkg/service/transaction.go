package service

import (
	"API/pkg/repository"
	"errors"
	"sync"
)

type TransactionService struct {
	repos repository.Transaction
}

func NewTransactionService(repos repository.Transaction) *TransactionService {
	return &TransactionService{repos: repos}
}

func (s *TransactionService) MakeTransaction(senderUsername string, RecipientUsername string, amount float64, a Action) error {

	var wg sync.WaitGroup
	wg.Add(2)

	var senderID uint32
	var senderErr error

	go func() {
		senderID, senderErr = a.GetUserIDByUsername(senderUsername)
		wg.Done()
	}()

	var recipientID uint32
	var recipientErr error

	go func() {
		recipientID, recipientErr = a.GetUserIDByUsername(RecipientUsername)
		wg.Done()
	}()

	wg.Wait()

	if senderErr != nil {
		return senderErr
	}
	if recipientErr != nil {
		return recipientErr
	}

	wg.Add(1)

	var balance float64
	var balanceErr error

	go func() {
		balance, balanceErr = a.GetUserBalance(senderID)
		wg.Done()
	}()

	wg.Wait()

	if balanceErr != nil {
		return balanceErr
	}

	if amount > balance {
		return errors.New("there are not enough coins to make a transaction")
	}

	if err := s.repos.MakeTransaction(senderID, recipientID, amount); err != nil {
		return err
	}

	return nil
}
