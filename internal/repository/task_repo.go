package repository

import (
	"todo-app/internal/domain"

	"github.com/jmoiron/sqlx"
)

type TaskRepository interface {
	CreateTask(task domain.Task) error
	DeleteById(id int) error
	UpdateById(id int) error
	GetAllTasks(usr int) ([]domain.Task, error)
	GetAllPendingTasks(usr int) ([]domain.Task, error)
}

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task domain.Task) error {
	_, err := r.db.Exec("INSERT INTO Task (title, description, completed, deleted, user_id, category_id) VALUES ($1, $2, $3, $4, $5, $6)", task.Title, task.Description, task.Completed, task.Deleted, task.User, task.Category)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteById(id int) error {
	_, err := r.db.Exec("UPDATE task SET deleted = true WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) UpdateById(id int) error {
	_, err := r.db.Exec("UPDATE task SET completed = true WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetAllPendingTasks(usr int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1 AND completed = false", usr)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetAllTasks(usr int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1", usr)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
