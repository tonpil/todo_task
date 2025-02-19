package usecase

import (
	"context"
	"ifigurin12/todo/interactor/dto"
	"ifigurin12/todo/interactor/iface/repository"
)

type GetTasksUseCase struct {
	TaskRepository repository.TaskRepository
}

func (u *GetTasksUseCase) Execute(ctx context.Context, inputDTO dto.GetTasksInputDTO) (*dto.GetTasksOutputDTO, error) {
	tasks, err := u.TaskRepository.ListTasks(ctx)

	if err != nil {
		return nil, err
	}

	return &dto.GetTasksOutputDTO{
		Items: tasks,
	}, nil
}
