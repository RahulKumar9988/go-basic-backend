package config

import (
<<<<<<< HEAD
	"log"
=======
	"fmt"
	// "log"
>>>>>>> 78f72f6493074c6e33cd9e2e109d32cec8e4be68
	"os"
)

type Config struct {
<<<<<<< HEAD
	DBHost string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBUser string `yaml:"DB_HOST" env:"DB_HOST" env-required:"true"`
	DBPass string `yaml:"DB_PASSWORD" env:"DB_PASSWORD" env-required:"true"`
	DBName string `yaml:"DB_NAME" env:"DB_NAME" env-required:"true"`
	DBPort string `yaml:"DB_PORT" env:"DB_PORT" env-required:"true"`
=======
	DBHost string
	DBUser string
	DBName string
	DBPass string
	DBPort string
>>>>>>> 78f72f6493074c6e33cd9e2e109d32cec8e4be68
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
<<<<<<< HEAD
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_Name"),
		DBPort: os.Getenv("DB_Port"),
	}

	if cfg.DBHost == "" || cfg.DBName == "" || cfg.DBPass == "" || cfg.DBPort == "" || cfg.DBUser == "" {
		log.Fatal(" missing env variable")
	}

	return cfg

=======
		DBName: os.Getenv("DB_NAME"),
		DBPass: os.Getenv("DB_PASS"),
		DBPort: os.Getenv("DB_PORT"),
	}

	if cfg.DBHost == "" || cfg.DBName == "" || cfg.DBUser == "" || cfg.DBPass == "" || cfg.DBPort == "" {
		// log.Fatal("DB CONNECTGION FAILED: MISSING ENV")
		fmt.Println("DB CONNECTGION FAILED: MISSING ENV")
	}

	return cfg
>>>>>>> 78f72f6493074c6e33cd9e2e109d32cec8e4be68
}
