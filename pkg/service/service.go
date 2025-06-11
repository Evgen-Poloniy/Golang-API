package service

import "API/pkg/repository"

type Authorization interface {
	CreateUser(*repository.Users) (int, error)
}

type Action interface {
	GetUserById(user_id int) (*repository.Users, error)
	GetUserByUsername(username string) (*repository.Users, error)
	GetUserIdByUsername(username string) (uint32, error)
	GetUserByAttributes(attributes map[string]string) (*repository.Users, error)
	GetUserBalance(user_id uint32) (float64, error)
}

type Transaction interface {
	MakeTransaction(senderUsername string, recipientUsername string, amount float64, a Action) error
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
