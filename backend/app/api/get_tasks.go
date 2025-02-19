package api

import entity "ifigurin12/todo/domain/task"

type GetTasksResponse struct {
	Result []entity.Task `json:"result" validate:"required"`
}
