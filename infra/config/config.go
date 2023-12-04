package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiRestPort int `envconfig:"API_REST_PORT"`
	EnableCron  bool `envconfig:"ENABLE_CRON"`

}

var Env Config = Config{
	ApiRestPort: 8080,
	EnableCron:  false,
}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := envconfig.Process("", &Env); err != nil {
		panic(err)
	}

	return nil
}
