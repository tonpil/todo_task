package main

import (
	"context"
	"ifigurin12/todo/app"
	"ifigurin12/todo/config"
)

// @title Todo Task API
// @version 1.0
// @description API for managing todo tasks
// @host localhost:8040
// @BasePath /
func main() {
	ctx := context.Background()
	postgresPool, err := config.NewPostgresPool(ctx)
	if err != nil {
		panic("could not create postgres pool")
	}

	err = postgresPool.Ping(ctx)

	if err != nil {
		panic("could not ping postgres")
	}

	app := app.Serve(ctx, postgresPool)

	defer app.Shutdown()
	app.Listen(":8040")
}
