package repository

import (
	"todo-app/internal/domain"

	"github.com/jmoiron/sqlx"
)

type VerificationRepository interface {
	Save(ev *domain.EmailVerification) error
	FindByToken(token string) (*domain.EmailVerification, error)
	MarkAsUsed(token string) error
	CheckIfAuth(usr int) bool
}

type verificationRepository struct {
	db *sqlx.DB
}

func NewVerificationRepository(db *sqlx.DB) VerificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) Save(ev *domain.EmailVerification) error {
	_, err := r.db.Exec(`INSERT INTO Email_verification VALUES ($1, $2, $3, $4)`, ev.Token, ev.UserId, ev.ExpiresAt, ev.Used)

	if err != nil {
		return err
	}

	return nil
}

func (r *verificationRepository) FindByToken(token string) (*domain.EmailVerification, error) {
	ev := domain.EmailVerification{}

	err := r.db.Select(&ev, "SELECT * FROM Email_verification WHERE token = $1", token)

	if err != nil {
		return nil, err
	}

	return &ev, nil
}

func (r *verificationRepository) MarkAsUsed(token string) error {
	_, err := r.db.Exec("UPDATE Email_verification SET used = true WHERE token = $1", token)

	if err != nil {
		return err
	}

	return nil
}

func (r *verificationRepository) CheckIfAuth(usr int) bool {
	var res int

	err := r.db.Get(&res, "SELECT COUNT(*) FROM email_verification WHERE user_id = $1 AND used = true", usr)

	if err != nil {
		return false
	}

	return res == 1

}
