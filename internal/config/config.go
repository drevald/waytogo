package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"os"
)

type Config struct {
	//Port uint16 `env:"PORT" default:"8080"`
	Port     uint16 `env:"PORT,required"`
	LogLevel string `LOG_LEVEL:"debug" default:"debug"`
	DbUrl    string `env:"DB_URL,required"`
}

func New(di *do.Injector) (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load()
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
