package dto

import entity "ifigurin12/todo/domain/task"

type GetTasksInputDTO struct {
}

type GetTasksOutputDTO struct {
	Items []entity.Task
}
