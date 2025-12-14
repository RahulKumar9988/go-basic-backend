package config

import (
	"fmt"
	// "log"
	"os"
)

type Config struct {
	DBHost string
	DBUser string
	DBName string
	DBPass string
	DBPort string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		DBPass: os.Getenv("DB_PASS"),
		DBPort: os.Getenv("DB_PORT"),
	}

	if cfg.DBHost == "" || cfg.DBName == "" || cfg.DBUser == "" || cfg.DBPass == "" || cfg.DBPort == "" {
		// log.Fatal("DB CONNECTGION FAILED: MISSING ENV")
		fmt.Println("DB CONNECTGION FAILED: MISSING ENV")
	}

	return cfg
}
