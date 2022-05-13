package balance

import (
	"github.com/gofiber/fiber/v2"
	"restapi/pkg/common/models"
)

type TransferBalanceRequestBody struct {
	Amount     int `json:"amount"`
	FromUserId int `json:"fromUserId"`
	ToUserId   int `json:"toUserId"`
}

func (h handler) TransferBalance(c *fiber.Ctx) error {
	body := TransferBalanceRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if body.Amount == 0 {
		return fiber.NewError(fiber.StatusNotFound, "It is impossible to transfer zero funds")
	}
	var balanceFrom, balanceTo models.Balance
	var firstUserBalance, secondUserBalance Balance
	find := h.DB.First(&balanceFrom, body.FromUserId)
	find.Scan(&firstUserBalance)
	find = h.DB.First(&balanceTo, body.ToUserId)
	find.Scan(&secondUserBalance)
	if firstUserBalance.Amount-body.Amount < 0 {
		return fiber.NewError(fiber.StatusNotFound, "There is not enough money in the account to transfer funds")
	}
	balanceFrom.Amount = firstUserBalance.Amount - body.Amount
	balanceTo.Amount = secondUserBalance.Amount + body.Amount
	h.DB.Save(&balanceFrom)
	h.DB.Save(&balanceTo)
	return c.Status(fiber.StatusOK).JSON(&balanceTo)
}
