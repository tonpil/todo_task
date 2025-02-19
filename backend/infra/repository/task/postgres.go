package task

import (
	"context"
	"errors"
	entity "ifigurin12/todo/domain/task"
	"ifigurin12/todo/infra/repository/db"
	"ifigurin12/todo/infra/repository/private"
	"ifigurin12/todo/interactor/iface/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskPostgresRepository struct {
	postgresPool *pgxpool.Pool
}

func NewTaskPostgresRepository(postgresPool *pgxpool.Pool) *TaskPostgresRepository {
	return &TaskPostgresRepository{postgresPool}
}

func (r *TaskPostgresRepository) ListTasks(ctx context.Context) ([]entity.Task, error) {
	queries := db.New(r.postgresPool)

	items, err := queries.ListTasks(ctx)

	if err != nil {
		return nil, private.TransformError(err)
	}

	result := make([]entity.Task, len(items))

	for i, item := range items {
		var status *entity.Status

		if item.Status != nil {
			value := entity.CreateStatusFromString(*item.Status)

			if value != nil {
				status = value
			} else {
				return nil, errors.New("invalid status from db")
			}
		}

		result[i] = entity.Task{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Status:      status,
			CreatedAt:   &item.CreatedAt.Time,
			UpdatedAt:   &item.UpdatedAt.Time,
		}
	}

	return result, nil
}

func (r *TaskPostgresRepository) GetTaskByID(ctx context.Context, ID int32) (*entity.Task, error) {
	queries := db.New(r.postgresPool)

	item, err := queries.GetTaskByID(ctx, ID)

	if err != nil {
		return nil, private.TransformError(err)
	}

	var status *entity.Status

	if item.Status != nil {
		value := entity.CreateStatusFromString(*item.Status)

		if value != nil {
			status = value
		} else {
			return nil, errors.New("invalid status from db")
		}
	}

	result := entity.Task{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Status:      status,
		CreatedAt:   &item.CreatedAt.Time,
		UpdatedAt:   &item.UpdatedAt.Time,
	}

	return &result, nil
}

func (r *TaskPostgresRepository) CreateTask(ctx context.Context, args repository.CreateTaskArgs) (*int32, error) {
	queries := db.New(r.postgresPool)

	ID, err := queries.CreateTask(ctx, db.CreateTaskParams{
		Title:       args.Title,
		Description: args.Description,
	})

	if err != nil {
		return nil, private.TransformError(err)
	}

	return &ID, nil
}

func (r *TaskPostgresRepository) UpdateTask(ctx context.Context, ID int32, args repository.UpdateTaskArgs) error {
	queries := db.New(r.postgresPool)

	var status *string

	if args.Status != nil {
		temp := string(*args.Status)
		status = &temp
	}

	err := queries.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          ID,
		Title:       args.Title,
		Description: args.Description,
		Status:      status,
		UpdatedAt: pgtype.Timestamp{
			Time:  args.UpdatedAt,
			Valid: true,
		},
	})

	return private.TransformError(err)
}

func (r *TaskPostgresRepository) DeleteTask(ctx context.Context, ID int32) error {
	queries := db.New(r.postgresPool)

	err := queries.DeleteTask(ctx, ID)

	return private.TransformError(err)
}
