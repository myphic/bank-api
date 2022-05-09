package balance

import (
	"context"
	"github.com/go-redis/cache/v8"
	"restapi/pkg/common/models"
	appredis "restapi/pkg/common/redis"
	"restapi/pkg/logs"
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
func (h handler) GetBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var balance models.Balance

	if result := h.DB.First(&balance, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	key := id
	obj := balance
	if err := appredis.GetCache().Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		h.logger.Fatalln("Cannot set redis cache", err)
	}
	h.logger.Infoln("Add to cache", obj)
	return c.Status(fiber.StatusOK).JSON(&balance)
}
