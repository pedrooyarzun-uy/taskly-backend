package task

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Save(task Task) error
	GetById(id int) (*Task, error)
	GetAll() (*[]Task, error)
	DeleteById(id int) error
	UpdateById(id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Save(task Task) error {
	_, err := r.db.Exec("INSERT INTO task (title, description, completed, deleted, created_at) VALUES ($1, $2, $3, $4, NOW())", task.Title, task.Description, task.Completed, task.Deleted)

	return err
}

func (r *repository) GetById(id int) (*Task, error) {
	task := Task{}
	err := r.db.Get(&task, "SELECT * FROM task WHERE id = $1", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &task, err
}

func (r *repository) GetAll() (*[]Task, error) {
	tasks := []Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task")

	if err != nil {
		return nil, err
	}

	return &tasks, nil

}

func (r *repository) DeleteById(id int) error {
	_, err := r.db.Exec("UPDATE task SET deleted = 1 WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateById(id int) error {
	_, err := r.db.Exec("UPDATE task SET completed = 1 WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
