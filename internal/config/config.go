package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/samber/do"	
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	//Port uint16 `env:"PORT" default:"8080"`
	Port uint16 `env:"PORT,required"`
	LogLevel string `LOG_LEVEL:"debug" default:"error"`
}

func New(di *do.Injector) (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load()
	}
	var cfg Config
	if err:=env.Parse(&cfg); err!= nil {
		return nil, err
	}
	return &cfg, nil
}
