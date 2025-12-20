package config

import (
	"errors"
	"os"
)

type Config struct {
	DBHost string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBUser string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBPass string `yaml:"DB_PASSWORD" env:"DB_PASSWORD" env-required:"true"`
	DBName string `yaml:"DB_NAME" env:"DB_NAME" env-required:"true"`
	DBPort string `yaml:"DB_PORT" env:"DB_PORT" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
	}

	// validate required env variables
	if err := Validation(cfg); err != nil {
		return nil, err
	}
	return cfg, nil

}

func Validation(cfg *Config) error {
	if cfg.DBHost == "" {
		return errors.New("DB_HOST missing")
	}
	if cfg.DBName == "" {
		return errors.New("DB_NAME is missing")
	}
	if cfg.DBPass == "" {
		return errors.New("DB_PASSWORD is missing")
	}
	if cfg.DBPort == "" {
		return errors.New("DB_PORT is missing")
	}
	if cfg.DBUser == "" {
		return errors.New("DB_USER is missing")
	}
	return nil
}
