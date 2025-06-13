package service

import (
	"API/pkg/repository"
	"errors"
	"sync"
)

type TransactionService struct {
	repos  repository.Transaction
	action Action
}

func NewTransactionService(repos repository.Transaction, a Action) *TransactionService {
	return &TransactionService{
		repos:  repos,
		action: a,
	}
}

func (s *TransactionService) MakeTransaction(senderUsername string, recipientUsername string, amount float64) (uint32, uint32, error) {

	var wg sync.WaitGroup
	wg.Add(2)

	var senderID uint32
	var senderErr error

	go func() {
		senderID, senderErr = s.action.GetUserIDByUsername(senderUsername)
		wg.Done()
	}()

	var recipientID uint32
	var recipientErr error

	go func() {
		recipientID, recipientErr = s.action.GetUserIDByUsername(recipientUsername)
		wg.Done()
	}()

	wg.Wait()

	if senderErr != nil {
		return 0, 0, senderErr
	}
	if recipientErr != nil {
		return 0, 0, recipientErr
	}

	wg.Add(1)

	var balance float64
	var balanceErr error

	go func() {
		balance, balanceErr = s.action.GetUserBalance(senderID)
		wg.Done()
	}()

	wg.Wait()

	if balanceErr != nil {
		return 0, 0, balanceErr
	}

	if amount > balance {
		return 0, 0, errors.New("there are not enough coins to make a transaction")
	}

	if err := s.repos.MakeTransaction(senderID, recipientID, amount); err != nil {
		return 0, 0, err
	}

	return senderID, recipientID, nil
}

func (s *TransactionService) CreateRecordOfTransaction(transaction *repository.Transactions) (uint32, error) {
	return s.repos.CreateRecordOfTransaction(transaction)
}

func (s *TransactionService) GetTransactionByID(transaction_id uint32) (*repository.Transactions, error) {
	return s.repos.GetTransactionByID(transaction_id)
}
