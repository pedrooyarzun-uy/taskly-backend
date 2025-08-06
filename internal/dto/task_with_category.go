package dto

import "time"

type TaskWithCategory struct {
	Id           int       `db:"id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	User         int       `db:"user_id"`
	Completed    bool      `db:"completed"`
	Deleted      bool      `db:"deleted"`
	Category     int       `db:"category_id"`
	CategoryName string    `db:"category_name"`
	CreatedAt    time.Time `db:"created_at"`
}
