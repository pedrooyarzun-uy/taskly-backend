package repository

import (
	"database/sql"
	"todo-app/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	UserExists(email string) bool
	GetUserByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user domain.User) error {

	_, err := r.db.Exec(`INSERT INTO Users (name, email, password, deleted, created_at) VALUES ($1,$2,$3,$4,$5)`, user.Name, user.Email, user.Password, false, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil

}

func (r *userRepository) UserExists(email string) bool {
	var res int

	err := r.db.Get(&res, "SELECT COUNT(*) FROM users WHERE email = $1", email)

	if err != nil {
		return false
	}

	return res == 1
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
