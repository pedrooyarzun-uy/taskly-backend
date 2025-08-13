package service

import (
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
)

type TaskService interface {
	CreateTask(task dto.CreateTaskRequest, userId int) error
	CompleteTask(task dto.CompleteTaskRequest, userId int) error
	IncompleteTask(task dto.CompleteTaskRequest, userId int) error
	DeleteTask(task dto.DeleteTaskRequest, userId int) error
	ModifyTask(task dto.ModifyTaskRequest, userId int) error
	GetAllTasks(userId int) ([]domain.Task, error)
	GetPendingTasks(userId int) ([]dto.CategoryWithTasks, error)
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

func (s *taskService) CompleteTask(task dto.CompleteTaskRequest, userId int) error {
	err := s.tr.CompleteTask(task.Id, userId)

	return err
}

func (s *taskService) IncompleteTask(task dto.CompleteTaskRequest, userId int) error {
	return s.tr.IncompleteTask(task.Id, userId)
}

func (s *taskService) DeleteTask(task dto.DeleteTaskRequest, userId int) error {
	err := s.tr.DeleteById(task.Id, userId)

	return err
}

func (s *taskService) ModifyTask(t dto.ModifyTaskRequest, userId int) error {
	task, err := s.tr.GetById(t.Id)

	if err != nil {
		return err
	}

	if t.Title != "" {
		task.Title = t.Title
	}

	if t.Description != "" {
		task.Description = t.Description
	}

	if t.Category != 0 {
		task.Category = t.Category
	}

	err = s.tr.ModifyById(task, userId)

	return err

}

func (s *taskService) GetAllTasks(userId int) ([]domain.Task, error) {
	return s.tr.GetAllTasks(userId)
}

func (s *taskService) GetPendingTasks(userId int) ([]dto.CategoryWithTasks, error) {
	tasks, err := s.tr.GetPendingTasks(userId)

	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]*dto.CategoryWithTasks)

	for _, row := range tasks {
		if _, ok := categoryMap[row.Category]; !ok {
			categoryMap[row.Category] = &dto.CategoryWithTasks{
				CategoryID:   row.Category,
				CategoryName: row.CategoryName,
				Tasks:        []dto.TaskResponse{},
			}
		}

		task := dto.TaskResponse{
			ID:          row.Id,
			Title:       row.Title,
			Description: row.Description,
			Completed:   row.Completed,
			Deleted:     row.Deleted,
		}

		categoryMap[row.Category].Tasks = append(categoryMap[row.Category].Tasks, task)
	}

	// Convertir el map a slice
	result := make([]dto.CategoryWithTasks, 0, len(categoryMap))
	for _, cat := range categoryMap {
		result = append(result, *cat)
	}

	return result, nil

}
