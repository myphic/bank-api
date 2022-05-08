package balance

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"restapi/pkg/logs"
)

const apiUrl = "api"

type handler struct {
	DB     *gorm.DB
	logger *logs.Logger
}

func RegisterRoutes(app *fiber.App, db *gorm.DB, logger *logs.Logger) {
	h := &handler{
		DB:     db,
		logger: logger,
	}

	routes := app.Group("/")
	routes.Post(apiUrl+"/deposit/", h.AddBalance)
	routes.Get(apiUrl+"/balance/:id", h.GetBalance)
	routes.Post(apiUrl+"/transfer/", h.TransferBalance)
	routes.Delete(apiUrl+"/decrease/:id/:amount", h.DeleteBalance)
}
