package dto

import entity "ifigurin12/todo/domain/task"

type UpdateTaskInputDTO struct {
	ID          int32
	Title       string
	Description *string
	Status      *entity.Status
}

type UpdateTaskOutputDTO struct {
}
