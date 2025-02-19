package controller

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/app/controller/private"
	"ifigurin12/todo/app/presenter"
	repository "ifigurin12/todo/infra/repository/task"
	"ifigurin12/todo/interactor/domain/task"
	"ifigurin12/todo/interactor/dto"
	usecase "ifigurin12/todo/interactor/use_case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body api.UpdateTaskRequest true "Updated task data"
// @Success 200 {object} api.UpdateTaskResponse "Task updated successfully"
// @Router /tasks/{id} [put]
func (s *TaskController) UpdateTask(c *fiber.Ctx) error {
	IDStr := c.Params("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var request api.UpdateTaskRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var statusFromRequest *task.Status

	if request.Status != nil {
		statusFromRequest = task.CreateStatusFromString(*request.Status)
		if statusFromRequest == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status from body"})
		}
	}

	presenter := presenter.UpdateTaskPresenter{}
	inputDTO := dto.UpdateTaskInputDTO{
		ID:          int32(ID),
		Title:       request.Title,
		Description: request.Description,
		Status:      statusFromRequest,
	}

	taskRepository := repository.NewTaskPostgresRepository(s.postgresPool)

	useCase := usecase.UpdateTaskUseCase{
		TaskRepository: taskRepository,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, errMsg := private.TransformErrorToHttpError(err)
		return c.Status(code).JSON(fiber.Map{"error": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(presenter.Present(outputDTO))
}
