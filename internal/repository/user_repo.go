package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-app/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetAllPendingTasks(id int) ([]domain.Task, error)
	GetAllTasks(id int) ([]domain.Task, error)
	UserExists(email string) bool
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user domain.User) error {

	_, err := r.db.Exec(`INSERT INTO Users (name, email, password, deleted, created_at) VALUES ($1,$2,$3,$4,$5)`, user.Name, user.Email, user.Password, true, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil

}

func (r *userRepository) GetAllPendingTasks(userId int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1 AND completed = false", userId)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *userRepository) GetAllTasks(userId int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1", userId)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *userRepository) UserExists(email string) bool {
	var res int

	err := r.db.Get(&res, "SELECT COUNT(*) FROM users WHERE email = $1", email)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return res == 1
}
