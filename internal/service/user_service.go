package service

import (
	"time"
	"todo-app/internal/domain"
	"todo-app/internal/repository"
)

type UserService interface {
	CreateUser(name string, email string, password string, createdAt time.Time) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(name string, email string, password string, createAt time.Time) error {
	user := domain.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createAt,
	}

	err := s.repo.CreateUser(user)

	return err
}
