package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
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
	ur repository.UserRepository
	vr repository.VerificationRepository
}

func NewUserService(
	ur repository.UserRepository,
	vr repository.VerificationRepository,
) UserService {
	return &userService{
		ur, vr,
	}
}

func (s *userService) CreateUser(usr dto.CreateUserRequest) error {
	res, err := s.ur.GetUserByEmail(usr.Email)

	if err != nil {
		fmt.Println(usr.Email)
		fmt.Println("ENTRE ACAAA")
		return err
	}

	if res != nil {
		err := errors.New("el usuario ya existe en el sistema")
		return err
	}

	password, _ := helpers.HashPassword(usr.Password)

	user := domain.User{
		Name:      usr.Name,
		Email:     usr.Email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err = s.ur.CreateUser(user)

	res, err = s.ur.GetUserByEmail(usr.Email)

	token := uuid.New().String()
	ev := domain.EmailVerification{
		Token:     token,
		UserId:    res.Id,
		ExpiresAt: time.Now(),
		Used:      false,
	}

	err = s.vr.Save(&ev)

	link := "http://localhost:8080/user/verify?token=" + token
	helpers.SendMail(user.Email, "Verifica tu cuenta!", "Necesitamos que verifiques tu cuenta! Ingresa a: "+link)

	return err
}
