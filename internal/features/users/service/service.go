package users_service

import (
	"fmt"

	"github.com/shitaiv1ck/todolist/internal/core/domains"
)

type UsersService struct {
	repository UsersRepository
}

type UsersRepository interface {
	CreateUser(user *domains.User) (*domains.User, error)
	FindByID(userID int) (*domains.User, error)
}

func NewService(repository UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (s *UsersService) CreateUser(user *domains.User) (*domains.User, error) {
	if err := user.EncryptePassword(); err != nil {
		return nil, fmt.Errorf("failed to encrypte password: %w", err)
	}

	createdUser, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return createdUser, err
}

func (s *UsersService) GetUser(userID int) (*domains.User, error) {
	foundUser, err := s.repository.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return foundUser, err
}
