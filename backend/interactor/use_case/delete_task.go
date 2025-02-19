package usecase

import (
	"context"
	"ifigurin12/todo/interactor/dto"
	"ifigurin12/todo/interactor/iface/repository"
)

type DeleteTaskUseCase struct {
	TaskRepository repository.TaskRepository
}

func (u *DeleteTaskUseCase) Execute(ctx context.Context, inputDTO dto.DeleteTaskInputDTO) (*dto.DeleteTaskOutputDTO, error) {
	_, err := u.TaskRepository.GetTaskByID(ctx, inputDTO.ID)

	if err != nil {
		return nil, err
	}

	err = u.TaskRepository.DeleteTask(
		ctx,
		inputDTO.ID,
	)

	if err != nil {
		return nil, err
	}

	return &dto.DeleteTaskOutputDTO{}, nil
}
