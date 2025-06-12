package service

import "API/pkg/repository"

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user *repository.Users) (uint32, error) {
	return s.repos.CreateUser(user)
}
