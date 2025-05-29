package service

import "API/pkg/repository"

type ActionService struct {
	repos repository.Action
}

func NewActionService(repos repository.Action) *ActionService {
	return &ActionService{repos: repos}
}

func (s *ActionService) GetUserById(user_id int) (*repository.Users, error) {
	return s.repos.GetUserById(user_id)
}
