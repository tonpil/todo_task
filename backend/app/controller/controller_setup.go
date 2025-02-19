package controller

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskController struct {
	ctx          context.Context
	postgresPool *pgxpool.Pool
}

func New(ctx context.Context, postgresPool *pgxpool.Pool) *TaskController {
	return &TaskController{ctx, postgresPool}
}
