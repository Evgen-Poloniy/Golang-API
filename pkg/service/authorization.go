package service

import (
	"API/pkg/attribute"
	"API/pkg/repository"
	"fmt"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(auth *attribute.AuthField) (uint32, error) {
	user := repository.Users{
		Username: auth.Username,
		Name:     auth.Name,
		Surname:  auth.Surname,
		Coins:    100,
	}

	var err error
	user.PasswordHash, err = encryptPassword(user.PasswordHash)
	if err != nil {
		return 0, fmt.Errorf("encryption failure: %v", err)
	}

	return s.repos.CreateUser(&user)
}

func (s *AuthService) GetPasswordHashByUsername(username string) (string, error) {
	return s.repos.GetPasswordHashByUsername(username)
}
