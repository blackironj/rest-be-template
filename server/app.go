package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	"github.com/blackironj/rest-be-template/env"
	"github.com/blackironj/rest-be-template/server/controller"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: env.SrvShutdownDeadline,
	})

	app.Use(recover.New())

	addHandler(app)

	return app
}

func addHandler(app *fiber.App) {
	if env.SrvEnv == "dev" {
		app.Get("/docs/*", swagger.HandlerDefault)
	}

	bookGroup := app.Group("books")

	bookGroup.Post("", controller.RegisterBook)
	bookGroup.Get("/:isbn", controller.GetBookByISBN)
}
