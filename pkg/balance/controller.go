package balance

import (
	"github.com/go-redis/cache/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"restapi/pkg/logs"
)

const apiUrl = "api"

type handler struct {
	DB     *gorm.DB
	logger *logs.Logger
	redis  *cache.Cache
}

func RegisterRoutes(app *fiber.App, db *gorm.DB, logger *logs.Logger, redis *cache.Cache) {
	h := &handler{
		DB:     db,
		logger: logger,
		redis:  redis,
	}

	routes := app.Group("/")
	routes.Post(apiUrl+"/deposit/", h.AddBalance)
	routes.Get(apiUrl+"/balance/:id", verifyCache, h.GetBalance)
	routes.Post(apiUrl+"/transfer/", h.TransferBalance)
	routes.Delete(apiUrl+"/decrease/:id/:amount", h.DeleteBalance)
}
