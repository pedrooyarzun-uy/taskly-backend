package dto

import "time"

type TaskResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
	Deleted   bool      `json:"deleted"`
}

type CategoryWithTasks struct {
	CategoryID   int            `json:"category_id"`
	CategoryName string         `json:"category_name"`
	Tasks        []TaskResponse `json:"tasks"`
}
