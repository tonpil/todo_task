package presenter

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/interactor/dto"
)

type UpdateTaskPresenter struct {
}

func (p *UpdateTaskPresenter) Present(outputDTO *dto.UpdateTaskOutputDTO) *api.UpdateTaskResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.UpdateTaskResponse{}

	return response
}
