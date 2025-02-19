package dto

import entity "ifigurin12/todo/interactor/domain/task"

type GetTasksInputDTO struct {
}

type GetTasksOutputDTO struct {
	Items []entity.Task
}
