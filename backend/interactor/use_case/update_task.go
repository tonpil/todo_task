package usecase

import (
	"context"
	"ifigurin12/todo/interactor/dto"
	"ifigurin12/todo/interactor/iface/repository"
	"time"
)

type UpdateTaskUseCase struct {
	TaskRepository repository.TaskRepository
}

func (u *UpdateTaskUseCase) Execute(ctx context.Context, inputDTO dto.UpdateTaskInputDTO) (*dto.UpdateTaskOutputDTO, error) {
	_, err := u.TaskRepository.GetTaskByID(ctx, inputDTO.ID)

	if err != nil {
		return nil, err
	}

	err = u.TaskRepository.UpdateTask(
		ctx,
		inputDTO.ID,
		repository.UpdateTaskArgs{
			Title:       inputDTO.Title,
			Description: inputDTO.Description,
			Status:      inputDTO.Status,
			UpdatedAt:   time.Now(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.UpdateTaskOutputDTO{}, nil
}
