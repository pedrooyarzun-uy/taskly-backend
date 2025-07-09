package domain

import "time"

type Task struct {
	Id          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	User        int       `db:"user_id"`
	Completed   bool      `db:"completed"`
	Deleted     bool      `db:"deleted"`
	Category    int       `db:"category_id"`
	CreatedAt   time.Time `db:"created_at"`
}
