package service

import (
	"errors"
	"log"
	"os"
	"time"
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/helpers"
	"todo-app/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(usr dto.CreateUserRequest) error
	VerifyUser(token string) error
	SignIn(usr dto.SignInRequest) (string, error)
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
		return err
	}

	if res != nil {
		err := errors.New("User already exists")
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
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(5)),
		Used:      false,
	}

	err = s.vr.Save(&ev)

	link := "<a href='" + os.Getenv("ALLOWED_ORIGINS") + "/verify-account?token=" + token + "'>" + "Verify your account" + "</a>"

	subject := "Please verify your account!"
	body := "Hello " + user.Name + ", <br>" +
		"Thank you for registering with us! To complete your registration, please verify your account by clicking the link below: <br>" +
		link + "<br>" +
		"If you did not create an account with us, please ignore this email.<br>" +
		"Best regards,<br>" +
		"The Taskly Team"
	err = helpers.SendMail(user.Email, subject, body)

	if err != nil {
		log.Println("Error sending email:", err)
	}

	return err
}

func (s *userService) VerifyUser(token string) error {
	return s.vr.MarkAsUsed(token)
}

func (s *userService) SignIn(req dto.SignInRequest) (string, error) {
	usr, err := s.ur.GetUserByEmail(req.Email)

	if err != nil || usr == nil || !helpers.VerifyPassword(req.Password, usr.Password) {
		return "", errors.New("Incorrect email or password")
	}

	token, err := helpers.GenerateJWT(usr.Id, usr.Name)

	if err != nil {
		return "", err
	}

	return token, nil
}
