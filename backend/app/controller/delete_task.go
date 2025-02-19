package controller

import (
	"ifigurin12/todo/app/controller/private"
	"ifigurin12/todo/app/presenter"
	repository "ifigurin12/todo/infra/repository/task"
	"ifigurin12/todo/interactor/dto"
	usecase "ifigurin12/todo/interactor/use_case"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 200 {object} api.DeleteTaskResponse "Successful response"
// @Router /tasks/{id} [delete]
func (s *TaskController) DeleteTask(c *fiber.Ctx) error {
	IDStr := c.Params("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	presenter := presenter.DeleteTaskPresenter{}
	inputDTO := dto.DeleteTaskInputDTO{
		ID: int32(ID),
	}

	taskRepository := repository.NewTaskPostgresRepository(s.postgresPool)

	useCase := usecase.DeleteTaskUseCase{
		TaskRepository: taskRepository,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, errMsg := private.TransformErrorToHttpError(err)
		return c.Status(code).JSON(fiber.Map{"error": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(presenter.Present(outputDTO))
}
