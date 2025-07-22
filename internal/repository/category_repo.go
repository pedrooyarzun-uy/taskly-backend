package repository

import (
	"todo-app/internal/domain"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	CreateCategory(cat domain.Category) error
	DeleteCategory(cat domain.Category) error
	GetById(id int, userId int) (domain.Category, error)
	ModifyCategory(cat domain.Category) error
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(cat domain.Category) error {
	_, err := r.db.Exec("INSERT INTO category (name, user_id, deleted) VALUES ($1, $2, $3)", cat.Name, cat.User, cat.Deleted)

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) ModifyCategory(cat domain.Category) error {
	_, err := r.db.Exec("UPDATE category SET name = $1 WHERE id = $2 AND user_id = $3", cat.Name, cat.Id, cat.User)

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) DeleteCategory(cat domain.Category) error {
	_, err := r.db.Exec("UPDATE category SET deleted = true WHERE id = $1 AND user_id = $2", cat.Id, cat.User)

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) GetById(id int, userId int) (domain.Category, error) {
	var cat domain.Category

	err := r.db.Get(&cat, "SELECT * FROM category WHERE id = $1 AND user_id = $2", id, userId)

	if err != nil {

		return domain.Category{}, err
	}

	return cat, nil
}
