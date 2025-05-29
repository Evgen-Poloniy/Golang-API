package service

import "API/pkg/repository"

type Authorization interface {
	CreateUser(*repository.Users) (int, error)
}

type Action interface {
	GetUserById(user_id int) (*repository.Users, error)
}

type Transaction interface {
	MakeTransaction(idSender uint64, idRecipient uint64) error
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
	}
}
