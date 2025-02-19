package controller

import (
	"ifigurin12/todo/app/api"
	"ifigurin12/todo/app/controller/private"
	"ifigurin12/todo/app/presenter"
	repository "ifigurin12/todo/infra/repository/task"
	"ifigurin12/todo/interactor/dto"
	usecase "ifigurin12/todo/interactor/use_case"

	"github.com/gofiber/fiber/v2"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the provided title and description
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body api.CreateTaskRequest true "Task data"
// @Success 200 {object} api.CreateTaskResponse "Successful response"
// @Router /tasks [post]
func (s *TaskController) CreateTask(c *fiber.Ctx) error {
	var input api.CreateTaskRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	presenter := presenter.CreateTaskPresenter{}
	inputDTO := dto.CreateTaskInputDTO{
		Title:       input.Title,
		Description: input.Description,
	}

	taskRepository := repository.NewTaskPostgresRepository(s.postgresPool)

	useCase := usecase.CreateTaskUseCase{
		TaskRepository: taskRepository,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, errMsg := private.TransformErrorToHttpError(err)
		return c.Status(code).JSON(fiber.Map{"error": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(presenter.Present(outputDTO))
}
