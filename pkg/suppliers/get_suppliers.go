package suppliers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier

	if result := h.DB.Find(&suppliers); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&suppliers)
}
