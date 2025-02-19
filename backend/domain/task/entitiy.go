package task

import "time"

type Task struct {
	ID          int32      `json:"id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description *string    `json:"description"`
	Status      *Status    `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
