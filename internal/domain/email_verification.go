package domain

import "time"

type EmailVerification struct {
	Token     string    `db:"uuid"`
	UserId    int       `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
	Used      bool      `db:"used"`
}
