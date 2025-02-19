package presenter

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/interactor/dto"
)

type DeleteTaskPresenter struct {
}

func (p *DeleteTaskPresenter) Present(outputDTO *dto.DeleteTaskOutputDTO) *api.DeleteTaskResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.DeleteTaskResponse{}

	return response
}
