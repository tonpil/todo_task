package api

import entity "ifigurin12/todo/interactor/domain/task"

type GetTasksResponse struct {
	Result []entity.Task `json:"result" validate:"required"`
}
