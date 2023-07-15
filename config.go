package main

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Port         string `env:"SERVICE_PORT,required"`
	DBConnString string `env:"DB_CONN_STRING"`
}

func ReadConfig() (config, error) {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	log.SetFormatter(&log.JSONFormatter{})

	return cfg, nil
}
