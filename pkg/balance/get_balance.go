package balance

import (
	"restapi/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var balance models.Balance

	if result := h.DB.First(&balance, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&balance)
}
