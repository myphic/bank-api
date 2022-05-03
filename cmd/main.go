package main

import (
	"log"

	"restapi/pkg/balance"
	"restapi/pkg/common/config"
	"restapi/pkg/common/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	balance.RegisterRoutes(app, h)

	err = app.Listen(c.Port)
	if err != nil {
		log.Fatalln("Failed at listen", err)
	}
}
