package users

import (
    "go-fiber-api-docker/pkg/common/models"
    "github.com/gofiber/fiber/v2"
)

type AddUserRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"-"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
    body := AddUserRequestBody{}

    // parse body, attach to AddUserRequestBody struct
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Error parsing request body: "+err.Error())
    }

    var user models.User
	// Validate user structure
    if err := user.SetPassword(body.Password); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Error setting password: "+err.Error())
    }
	
    user.Email = body.Email
	
	if err := user.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
    if result := h.DB.Create(&user); result.Error != nil {
        // Puedes cambiar el código de estado según el tipo de error de la base de datos
        return fiber.NewError(fiber.StatusUnprocessableEntity, "Error creating user: "+result.Error.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(&user)
}
