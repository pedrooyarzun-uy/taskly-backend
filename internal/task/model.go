package task

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	Deleted     bool
	CreatedAt   time.Time `db:"created_at"`
}
