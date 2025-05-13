package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-app/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetAllPendingTasks(id int) ([]domain.Task, error)
	GetAllTasks(id int) ([]domain.Task, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user domain.User) error {
	return nil
}

func (r *userRepository) GetAllPendingTasks(userId int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1 AND completed = 0", userId)

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
