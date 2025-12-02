package config

import "os"

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost: os.Getenv("DB_Host"),
		DBUser: os.Getenv("DB_User"),
		DBPass: os.Getenv("DB_Pass"),
		DBName: os.Getenv("DB_Name"),
		DBPort: os.Getenv("DB_Port"),
	}
}
