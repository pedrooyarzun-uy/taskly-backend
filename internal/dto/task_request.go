package dto

import "time"

type CreateTaskRequest struct {
	Title    string    `json:"title" binding:"required"`
	DueDate  time.Time `json:"due_date" binding:"required"`
	Category int       `json:"category" binding:"required"`
}

type CompleteTaskRequest struct {
	Id int `json:"id" binding:"required"`
}

type DeleteTaskRequest struct {
	Id int `json:"id" binding:"required"`
}

type ModifyTaskRequest struct {
	Id       int    `json:"id" binding:"required"`
	Title    string `json:"title"`
	Category int    `json:"category"`
}
