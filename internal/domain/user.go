package domain

import "time"

type User struct {
	Id        int       `db:"int"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	Deleted   bool      `db:"deleted"`
}
