package service

import "API/pkg/repository"

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user *repository.Users) (int, error) {
	return s.repos.CreateUser(user)
}
