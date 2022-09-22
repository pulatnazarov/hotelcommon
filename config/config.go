package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/pulatnazarov/hotelcommon/postgres"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	postgres.Config
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
