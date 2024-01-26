package users

import (
	"go-fiber-api-docker/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequestBody struct {
	Email  string `json:"email"`
	Password string `json:"-"`
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateUserRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	if err := user.SetPassword(body.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user.Email = body.Email
	
	if err := user.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// save user
	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(&user)
}
