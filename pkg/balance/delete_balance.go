package balance

import (
	"gorm.io/gorm"
	"restapi/pkg/common/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteBalance(c *fiber.Ctx) error {
	getId := c.Params("id")
	amountFromParams := c.Params("amount")

	var balance models.Balance

	id, _ := strconv.Atoi(getId)
	balance.Id = id
	amount, _ := strconv.Atoi(amountFromParams)
	var currentBalance Balance
	if find := h.DB.First(&balance, id).Scan(&currentBalance); find.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, find.Error.Error())
	}

	if currentBalance.Amount-int32(amount) < 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "There is not enough money in the account to withdraw funds")
	}

	h.DB.Model(&balance).Update("amount", gorm.Expr("amount - ?", amountFromParams))

	return c.SendStatus(fiber.StatusOK)
}
