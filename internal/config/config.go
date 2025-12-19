package config

import (
	"log"
	"os"
)

type Config struct {
	DBHost string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBUser string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBPass string `yaml:"DB_PASSWORD" env:"DB_PASSWORD" env-required:"true"`
	DBName string `yaml:"DB_NAME" env:"DB_NAME" env-required:"true"`
	DBPort string `yaml:"DB_PORT" env:"DB_PORT" env-required:"true"`
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_Name"),
		DBPort: os.Getenv("DB_Port"),
	}

	if cfg.DBHost == "" || cfg.DBName == "" || cfg.DBPass == "" || cfg.DBPort == "" || cfg.DBUser == "" {
		log.Fatal(" missing env variable")
	}

	return cfg

}
