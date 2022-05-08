package config

import (
	"github.com/spf13/viper"
	"restapi/pkg/logs"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
}

func LoadConfig() (c Config, err error) {
	logger := logs.GetLogger()
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		logger.Fatalln("Cannot read config")
	}

	err = viper.Unmarshal(&c)

	return
}
