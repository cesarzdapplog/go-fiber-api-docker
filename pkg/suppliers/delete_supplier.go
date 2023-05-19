package suppliers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")

	var supplier models.Supplier

	if result := h.DB.First(&supplier, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete supplier from db
	h.DB.Delete(&supplier)

	return c.SendStatus(fiber.StatusOK)
}
