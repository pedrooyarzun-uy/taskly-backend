package domain

import "time"

type Task struct {
	Id        int       `db:"id"`
	Title     string    `db:"title"`
	User      int       `db:"user_id"`
	DueDate   time.Time `db:"due_date"`
	Completed bool      `db:"completed"`
	Deleted   bool      `db:"deleted"`
	Category  int       `db:"category_id"`
	CreatedAt time.Time `db:"created_at"`
}
