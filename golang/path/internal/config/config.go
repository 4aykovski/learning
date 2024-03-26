package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	STORAGE string `yaml:"storage_path"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(cfg.STORAGE); os.IsNotExist(err) {
		panic(err)
	}

	return &cfg, nil
}
