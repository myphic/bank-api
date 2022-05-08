package db

import (
	"fmt"
	"restapi/pkg/logs"

	"restapi/pkg/common/config"
	"restapi/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	logger := logs.GetLogger()
	if err != nil {
		logger.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Balance{})
	if err != nil {
		logger.Fatalln(err)
	}

	return db
}
