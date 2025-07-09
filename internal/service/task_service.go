package service

import (
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
)

type TaskService interface {
	CreateTask(task dto.CreateTaskRequest, userId int) error
}

type taskService struct {
	tr repository.TaskRepository
}

func NewTaskService(
	tr repository.TaskRepository,
) TaskService {
	return &taskService{
		tr,
	}
}

func (s *taskService) CreateTask(task dto.CreateTaskRequest, userId int) error {

	domTask := domain.Task{
		Title:       task.Title,
		Description: task.Description,
		User:        userId,
		Completed:   false,
		Deleted:     false,
		Category:    task.Category,
	}

	err := s.tr.CreateTask(domTask)

	return err

}
