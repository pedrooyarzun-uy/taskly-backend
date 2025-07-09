package dto

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}
