package service

import "API/pkg/repository"

type Authorization interface {
	CreateUser(*repository.Users) (int, error)
}

type Messenger interface {
	SendMessage(senderName *string, recipientName *string) error
}

type Transaction interface {
	MakeTransaction(idSender uint64, idRecipient uint64) error
}

type Service struct {
	Authorization
	Messenger
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
