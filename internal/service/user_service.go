package service

import (
	"errors"
	"time"
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/helpers"
	"todo-app/internal/repository"
)

type UserService interface {
	CreateUser(usr dto.CreateUserRequest) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(usr dto.CreateUserRequest) error {

	password, _ := helpers.HashPassword(usr.Password)

	user := domain.User{
		Name:      usr.Name,
		Email:     usr.Email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if s.repo.UserExists(usr.Email) {
		err := errors.New("el usuario ya existe en el sistema")
		return err
	}

	err := s.repo.CreateUser(user)

	return err
}
