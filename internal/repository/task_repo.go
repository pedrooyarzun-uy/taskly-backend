package repository

import (
	"database/sql"
	"todo-app/internal/domain"
	"todo-app/internal/dto"

	"github.com/jmoiron/sqlx"
)

type TaskRepository interface {
	CreateTask(task domain.Task) error
	DeleteById(id int, userId int) error
	CompleteTask(taskId int, userId int) error
	GetAllTasks(usr int) ([]domain.Task, error)
	GetPendingTasks(usr int) ([]dto.TaskWithCategory, error)
	GetById(id int) (domain.Task, error)
	ModifyById(task domain.Task, usr int) error
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

func (r *taskRepository) DeleteById(id int, userId int) error {
	_, err := r.db.Exec("UPDATE task SET deleted = true WHERE id = $1 AND user_id = $2", id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) CompleteTask(taskId int, userId int) error {
	_, err := r.db.Exec("UPDATE task SET completed = true WHERE id = $1 AND user_id = $2", taskId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetPendingTasks(usr int) ([]dto.TaskWithCategory, error) {
	tasks := []dto.TaskWithCategory{}

	err := r.db.Select(&tasks, `SELECT t.*, c.name AS category_name
		FROM task t 
		JOIN category c ON c.id = t.category_id
		WHERE t.user_id = $1 AND t.completed = false AND t.deleted = false`, usr)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetAllTasks(usr int) ([]domain.Task, error) {
	tasks := []domain.Task{}

	err := r.db.Select(&tasks, "SELECT * FROM task WHERE user_id = $1 AND deleted = false", usr)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetById(id int) (domain.Task, error) {
	var task domain.Task

	err := r.db.Get(&task, "SELECT * FROM task WHERE id = $1", id)

	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) ModifyById(task domain.Task, usr int) error {
	_, err := r.db.Exec("UPDATE task SET title = $1, description = $2, category_id = $3 WHERE id = $4 AND user_id = $5", task.Title, task.Description, task.Category, task.Id, usr)

	return err
}
