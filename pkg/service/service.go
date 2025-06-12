package service

import "API/pkg/repository"

type Authorization interface {
	CreateUser(*repository.Users) (uint32, error)
}

type Action interface {
	GetUserByID(user_id uint32) (*repository.Users, error)
	GetUserByUsername(username string) (*repository.Users, error)
	GetUserIDByUsername(username string) (uint32, error)
	GetUserByAttributes(attributes map[string]string) (*repository.Users, error)
	GetUserBalance(user_id uint32) (float64, error)
}

type Transaction interface {
	MakeTransaction(senderUsername string, recipientUsername string, amount float64, a Action) (uint32, uint32, error)
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
		Transaction:   NewTransactionService(repos.Transaction),
	}
}
