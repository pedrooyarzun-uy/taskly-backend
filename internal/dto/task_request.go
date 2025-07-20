package dto

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    int    `json:"category" binding:"required"`
}

type CompleteTaskRequest struct {
	Id int `json:"id" binding:"required"`
}

type DeleteTaskRequest struct {
	Id int `json:"id" binding:"required"`
}

type ModifyTaskRequest struct {
	Id          int    `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    int    `json:"category"`
}
