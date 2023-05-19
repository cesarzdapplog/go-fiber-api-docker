package suppliers

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddSupplierRequestBody struct {
	Name  string `json:"name"`
	Address string `json:"address`
}

func (h handler) AddSupplier(c *fiber.Ctx) error {
	body := AddSupplierRequestBody{}

	// parse body, attach to AddSupplierRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var supplier models.Supplier

	supplier.Name = body.Name
	supplier.Address = body.Address

	// insert new db entry
	if result := h.DB.Create(&supplier); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&supplier)
}
