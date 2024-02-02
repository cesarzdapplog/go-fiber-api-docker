package auth

import (
    "go-fiber-api-docker/pkg/common/models"
    "github.com/gofiber/fiber/v2"
    "fmt"
)

type LoginRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h handler) LoginUser(c *fiber.Ctx) error {
    // Parsear el cuerpo JSON de la solicitud
    var reqBody LoginRequestBody
    if err := c.BodyParser(&reqBody); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Error parsing request body: "+err.Error())
    }

    // Buscar el usuario en la base de datos por correo electrónico
    var user models.User
    if result := h.DB.Where("email = ?", reqBody.Email).First(&user); result.Error != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Usuario no encontrado")
    }

    if err := user.CheckPassword(string(reqBody.Password)); err != nil {
        fmt.Println(err)
        return fiber.NewError(fiber.StatusUnauthorized, "Credenciales inválidas")
    }

    // Si las credenciales son válidas, generar token JWT (debes implementar esta función)
    token, err := generateJWTToken(user)
    if err != nil {
        return err
    }

    // Devolver el token al cliente
    return c.JSON(fiber.Map{"token": token})
}
