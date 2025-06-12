package service

import "API/pkg/repository"

type ActionService struct {
	repos repository.Action
}

func NewActionService(repos repository.Action) *ActionService {
	return &ActionService{repos: repos}
}

func (s *ActionService) GetUserByID(user_id int) (*repository.Users, error) {
	return s.repos.GetUserByID(user_id)
}

func (s *ActionService) GetUserByUsername(username string) (*repository.Users, error) {
	return s.repos.GetUserByUsername(username)
}

func (s *ActionService) GetUserIDByUsername(username string) (uint32, error) {
	return 0, nil
}

func (s *ActionService) GetUserByAttributes(attributes map[string]string) (*repository.Users, error) {
	return s.repos.GetUserByAttributes(attributes)
}

func (s *ActionService) GetUserBalance(user_id uint32) (float64, error) {
	return s.repos.GetUserBalance(user_id)
}
