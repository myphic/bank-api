package balance

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/balance")
	routes.Post("/", h.AddBalance)
	routes.Get("/:id", h.GetBalance)
	routes.Put("/:id&:amount", h.AddBalance)
	routes.Delete("/:id/:amount", h.DeleteBalance)
}
