package service

import "API/pkg/repository"

func (s *Service) SignUp(user *repository.Users) error {
	return s.repos.TestInsertUser(user)
}

func (s *Service) SignIn(user *repository.Users) error {
	if user.Username == "" || user.Password == "" {
		return nil
	}

	password, err := s.repos.GetPasswordByUsername(&user.Username)
	if err != nil {
		return err
	}

	if password != user.Password {
		return err
	}

	return nil
}
