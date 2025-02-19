package presenter

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/interactor/dto"
)

type GetTasksPresenter struct {
}

func (p *GetTasksPresenter) Present(outputDTO *dto.GetTasksOutputDTO) *api.GetTasksResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.GetTasksResponse{
		Result: outputDTO.Items,
	}

	return response
}
