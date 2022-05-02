package balance

import (
	"github.com/gofiber/fiber/v2"
	"restapi/pkg/common/models"
)

type TransferBalanceRequestBody struct {
	Amount     int32 `json:"amount"`
	FromUserId int   `json:"from_UserId"`
	ToUserId   int   `json:"to_UserId"`
}

func (h handler) TransferBalance(c *fiber.Ctx) error {
	body := TransferBalanceRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var balanceFrom models.Balance
	var balanceTo models.Balance
	var firstUserBalance, secondUserBalance Balance
	find := h.DB.First(&balanceFrom, body.FromUserId)
	find.Scan(&firstUserBalance)
	find = h.DB.First(&balanceTo, body.ToUserId)
	find.Scan(&secondUserBalance)
	if firstUserBalance.Amount-body.Amount < 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "There is not enough money in the account to transfer funds")
	}
	balanceFrom.Amount = firstUserBalance.Amount - body.Amount
	balanceTo.Amount = secondUserBalance.Amount + body.Amount
	h.DB.Save(&balanceFrom)
	h.DB.Save(&balanceTo)
	return c.Status(fiber.StatusOK).JSON(&balanceTo)

}
