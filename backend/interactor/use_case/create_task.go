package usecase

import (
	"context"
	"ifigurin12/todo/interactor/dto"
	"ifigurin12/todo/interactor/iface/repository"
)

type CreateTaskUseCase struct {
	TaskRepository repository.TaskRepository
}

func (u *CreateTaskUseCase) Execute(ctx context.Context, inputDTO dto.CreateTaskInputDTO) (*dto.CreateTaskOutputDTO, error) {
	item, err := u.TaskRepository.CreateTask(
		ctx,
		repository.CreateTaskArgs{
			Title:       inputDTO.Title,
			Description: inputDTO.Description,
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.CreateTaskOutputDTO{
		ID: *item,
	}, nil
}
