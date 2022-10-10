package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/blackironj/rest-be-template/server/common"
	"github.com/blackironj/rest-be-template/server/service"
)

// GetUser is a function to get an user by email
// @Summary Get an user by email
// @Description Get an user by email
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponseHTTP{data=model.User}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /users [get]
func GetUser(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "wrong email",
		})
	}

	user := service.GetUserByEmail(email)
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(common.ResponseHTTP{
		Success: false,
		Data:    user,
		Message: "user not found",
	})
}
