package service

import "API/pkg/repository"

type Authorization interface {
	CreateAccount(user *repository.Users) error
	DeleteAccount() error
}

type Messenger interface {
	SendMessage(idSender uint64, idRecipient uint64) error
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

/*
func (s *Service) CreateAccount(user *repository.Users) error {
	if err := s.repos.InsertUser(user); err != nil {
		return err
	}
	return nil
}
*/
