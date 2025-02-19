package presenter

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/interactor/dto"
)

type CreateTaskPresenter struct {
}

func (p *CreateTaskPresenter) Present(outputDTO *dto.CreateTaskOutputDTO) *api.CreateTaskResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.CreateTaskResponse{
		Result: api.CreateTaskResult{
			ID: outputDTO.ID,
		},
	}

	return response
}
