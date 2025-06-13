package service

import (
	"API/pkg/attribute"
	"API/pkg/repository"
)

type Authorization interface {
	CreateUser(auth *attribute.AuthField) (uint32, error)
	GetPasswordHashByUsername(username string) (string, error)
}

type Action interface {
	GetUserByID(user_id uint32) (*attribute.ActionField, error)
	GetUserByUsername(username string) (*attribute.ActionField, error)
	GetUserIDByUsername(username string) (uint32, error)
	GetUserByAttributes(attributes map[string]string) (*attribute.ActionField, error)
	GetUserBalance(user_id uint32) (float64, error)
}

type Transaction interface {
	MakeTransaction(senderUsername string, recipientUsername string, amount float64) (uint32, uint32, error)
	CreateRecordOfTransaction(transaction *repository.Transactions) (uint32, error)
	GetTransactionByID(transaction_id uint32) (*repository.Transactions, error)
}

type Service struct {
	Authorization
	Action
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Action:        NewActionService(repos.Action),
		Transaction:   NewTransactionService(repos.Transaction, repos.Action),
	}
}
