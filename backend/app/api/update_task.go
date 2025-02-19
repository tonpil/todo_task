package api

type UpdateTaskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
	Status      *string `json:"status" validate:"oneof=new in_progress done"`
}

type UpdateTaskResponse struct {
}
