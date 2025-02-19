package dto

type CreateTaskInputDTO struct {
	Title       string
	Description *string
}

type CreateTaskOutputDTO struct {
	ID int32
}
