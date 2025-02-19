package repository

import (
	"context"
	entity "ifigurin12/todo/domain/task"
	"time"
)

type TaskRepository interface {
	ListTasks(ctx context.Context) ([]entity.Task, error)
	GetTaskByID(ctx context.Context, ID int32) (*entity.Task, error)
	CreateTask(ctx context.Context, args CreateTaskArgs) (*int32, error)
	UpdateTask(ctx context.Context, ID int32, args UpdateTaskArgs) error
	DeleteTask(ctx context.Context, ID int32) error
}

type CreateTaskArgs struct {
	Title       string
	Description *string
}

type UpdateTaskArgs struct {
	Title       string
	Description *string
	Status      *entity.Status
	UpdatedAt   time.Time
}
