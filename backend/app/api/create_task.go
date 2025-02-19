package api

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type CreateTaskResponse struct {
	Result CreateTaskResult `json:"result" validate:"required"`
}

type CreateTaskResult struct {
	ID int32 `json:"id" validate:"required"`
}
