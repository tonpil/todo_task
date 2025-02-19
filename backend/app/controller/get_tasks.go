package controller

import (
	"ifigurin12/todo/app/controller/private"
	"ifigurin12/todo/app/presenter"
	repository "ifigurin12/todo/infra/repository/task"
	"ifigurin12/todo/interactor/dto"
	usecase "ifigurin12/todo/interactor/use_case"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get all tasks
// @Description Retrieve a list of all tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} api.GetTasksResponse "Successful response"
// @Router /tasks [get]
func (s *TaskController) GetTasks(c *fiber.Ctx) error {
	presenter := presenter.GetTasksPresenter{}
	inputDTO := dto.GetTasksInputDTO{}
	taskRepository := repository.NewTaskPostgresRepository(s.postgresPool)

	useCase := usecase.GetTasksUseCase{
		TaskRepository: taskRepository,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, errMsg := private.TransformErrorToHttpError(err)
		return c.Status(code).JSON(fiber.Map{"error": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(presenter.Present(outputDTO))
}
