package service

import "API/pkg/repository"

type Authorization interface {
	SignUp(*repository.Users) error
	SignIn(*repository.Users) error
	DeleteAccount() error
}

type Messenger interface {
	SendMessage(senderName *string, recipientName *string) error
}

type Transaction interface {
	MakeTransaction(idSender uint64, idRecipient uint64) error
}

type Service struct {
	repos *repository.Repository
	Authorization
	Messenger
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}
