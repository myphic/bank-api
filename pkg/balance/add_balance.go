package balance

import (
	"github.com/gofiber/fiber/v2"
	"restapi/pkg/common/models"
)

type AddBalanceRequestBody struct {
	Amount int `json:"amount"`
	Id     int `json:"id"`
}
type Balance struct {
	Amount int
}

func (h handler) AddBalance(c *fiber.Ctx) error {
	body := AddBalanceRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var balance models.Balance

	find := h.DB.First(&balance, body.Id)

	if find.Error == nil {
		var currentBalance Balance
		find.Scan(&currentBalance)
		balance.Amount = body.Amount + currentBalance.Amount
		h.DB.Save(&balance)
		return c.Status(fiber.StatusOK).JSON(&balance)
	} else {
		balance.Amount = body.Amount
		balance.Id = body.Id
		if result := h.DB.Create(&balance); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
		return c.Status(fiber.StatusCreated).JSON(&balance)
	}
}
