package dto

import "ifigurin12/todo/interactor/domain/task"

type UpdateTaskInputDTO struct {
	ID          int32
	Title       string
	Description *string
	Status      *task.Status
}

type UpdateTaskOutputDTO struct {
}
