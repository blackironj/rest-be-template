package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/blackironj/rest-be-template/env"
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
	app.Get("/api/list", func(c *fiber.Ctx) error {
		fmt.Println("🥉 Last handler")
		return c.SendString("Hello, World 👋!")
	})
}
