package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/blackironj/rest-be-template/server/service"
)

func GetUser(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "wrong email",
		})
	}

	user := service.GetUserByEmail(email)
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
