package task

import "errors"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTask(title string, description string) error {
	if title == "" {
		return errors.New("El título se encuentra vacío")
	}

	task := Task{
		Title:       title,
		Description: description,
		Completed:   false,
		Deleted:     false,
	}

	err := s.repo.Save(task)

	return err
}

func (s *Service) CompleteTask(id int) error {
	err := s.repo.UpdateById(id)

	return err
}

func (s *Service) DeleteTask(id int) error {
	err := s.repo.DeleteById(id)

	return err
}

func (s *Service) GetPendingTasks() (*[]Task, error) {
	res, err := s.repo.GetAll()

	return res, err
}
