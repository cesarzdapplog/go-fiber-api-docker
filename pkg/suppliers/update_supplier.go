package suppliers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateSuppliertRequestBody struct {
	Name  string `json:"name"`
	Address string `json:"address`
}

func (h handler) UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateSuppliertRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var supplier models.Supplier

	if result := h.DB.First(&supplier, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	supplier.Name = body.Name
	supplier.Address = body.Address

	// save supplier
	h.DB.Save(&supplier)

	return c.Status(fiber.StatusOK).JSON(&supplier)
}
