package dto

type TaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Deleted     bool   `json:"deleted"`
}

type CategoryWithTasks struct {
	CategoryID   int            `json:"category_id"`
	CategoryName string         `json:"category_name"`
	Tasks        []TaskResponse `json:"tasks"`
}
