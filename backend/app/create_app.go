package app

import (
	"context"
	"ifigurin12/todo/app/controller"

	_ "ifigurin12/todo/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Serve(ctx context.Context, postgresPool *pgxpool.Pool) *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	taskController := controller.New(ctx, postgresPool)

	taskGroup := app.Group("/tasks")

	taskGroup.Get("", taskController.GetTasks)
	taskGroup.Post("", taskController.CreateTask)
	taskGroup.Put("/:id", taskController.UpdateTask)
	taskGroup.Delete("/:id", taskController.DeleteTask)

	return app
}
