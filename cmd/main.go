package main

import (
	"restapi/pkg/logs"

	"restapi/pkg/balance"
	"restapi/pkg/common/config"
	"restapi/pkg/common/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logs.Init()
	logger := logs.GetLogger()
	logger.Info("Loading config..")
	c, err := config.LoadConfig()
	if err != nil {
		logger.Fatalln("Failed at config", err)
	}
	logger.Info("The config is loaded")
	h := db.Init(&c)
	app := fiber.New()

	balance.RegisterRoutes(app, h)

	err = app.Listen(c.Port)
	if err != nil {
		logger.Fatalln("Failed at listen", err)
	}
}
