package balance

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const apiUrl = "api"

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/")
	routes.Post(apiUrl+"/deposit/", h.AddBalance)
	routes.Get(apiUrl+"/balance/:id", h.GetBalance)
	routes.Post(apiUrl+"/transfer/", h.TransferBalance)
	routes.Delete(apiUrl+"/decrease/:id/:amount", h.DeleteBalance)
}
