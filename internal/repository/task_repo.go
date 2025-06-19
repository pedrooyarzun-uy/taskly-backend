package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-app/internal/domain"
)

type TaskRepository interface {
	CreateTask(task domain.Task) error
	DeleteById(id int) error
	UpdateById(id int) error
}

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task domain.Task) error {
	_, err := r.db.Exec("INSERT INTO task (title, description, completed, deleted, user) VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.Completed, task.Deleted, task.User)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteById(id int) error {
	_, err := r.db.Exec("UPDATE task SET deleted = 1 WHERE id = $1", id)

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
