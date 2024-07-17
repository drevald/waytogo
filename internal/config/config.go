package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port uint16 `env:"PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	var cfg Config
	if err:=env.Parse(&cfg); err!= nil {
		return nil, err
	}
	return &cfg, nil
}
