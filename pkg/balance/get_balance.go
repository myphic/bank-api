package balance

import (
	"context"
	"github.com/go-redis/cache/v8"
	"restapi/pkg/common/models"
	appredis "restapi/pkg/common/redis"
	"restapi/pkg/logs"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.TODO()

func verifyCache(c *fiber.Ctx) error {
	logger := logs.GetLogger()

	id := c.Params("id")
	var wanted Balance

	if err := appredis.GetCache().Get(ctx, id, &wanted); err == nil {
		logger.Infoln("Current cache:", wanted)
	} else {
		return c.Next()
	}
	data := &models.Balance{}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Cached": data})
}

func convertCurrency(currency float64, currencyFrom string) float64 {
	logger := logs.GetLogger()
	agent := fiber.Get("https://api.coingate.com/v2/rates/merchant/RUB/" + currencyFrom + "/")
	_, response, err := agent.String()
	if err != nil {
		logger.Errorln("Error with convert currency: ", err)
	}
	convertedCurrency, errConv := strconv.ParseFloat(response, 64)
	if errConv != nil {
		logger.Errorln("Error with convert string to int: ", err)
	}
	return convertedCurrency * currency
}

func (h handler) GetBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	currency := c.Params("currency")
	var balance models.Balance
	if result := h.DB.First(&balance, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	if len(currency) > 0 {
		balance.Amount = convertCurrency(balance.Amount, currency)
		return c.Status(fiber.StatusOK).JSON(&balance)
	}
	key := id
	obj := balance
	err := appredis.GetCache().Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	})
	if err != nil {
		h.logger.Fatalln("Cannot set redis cache", err)
	}
	h.logger.Infoln("Add to cache", obj)
	return c.Status(fiber.StatusOK).JSON(&balance)
}
