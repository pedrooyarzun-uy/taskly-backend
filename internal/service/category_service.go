package service

import (
	"todo-app/internal/domain"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest, userId int) error
	DeleteCategory(req dto.DeleteCategoryRequest, userId int) error
	ModifyCategory(req dto.ModifyCategoryRequest, userId int) error
	GetCategories(userId int) ([]domain.Category, error)
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

func (s *categoryService) DeleteCategory(req dto.DeleteCategoryRequest, userId int) error {
	cat := domain.Category{
		Id:   req.Id,
		User: userId,
	}

	return s.cr.DeleteCategory(cat)
}

func (s *categoryService) ModifyCategory(req dto.ModifyCategoryRequest, userId int) error {
	cat, err := s.cr.GetById(req.Id, userId)

	if err != nil {
		return err
	}

	cat.Name = req.Name

	return s.cr.ModifyCategory(cat)

}

func (s *categoryService) GetCategories(userId int) ([]domain.Category, error) {
	return s.cr.GetAllCategories(userId)
}
