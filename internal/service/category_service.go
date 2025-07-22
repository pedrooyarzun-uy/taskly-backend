package service

import (
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest, userId int) error
}

type categoryService struct {
	cr repository.CategoryRepository
}

func NewCategoryService(
	cr repository.CategoryRepository,
) CategoryService {
	return &categoryService{
		cr,
	}
}

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest, userId int) error {
	cat := domain.Category{
		Name:    req.Name,
		User:    userId,
		Deleted: false,
	}

	return s.cr.CreateCategory(cat)
}
