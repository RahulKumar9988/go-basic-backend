package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	ConnectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBHost,
		cfg.DBName,
		cfg.DBPass,
		cfg.DBPort,
		cfg.DBUser,
	)

	db, err := sql.Open("mysql", ConnectionString)

	if err != nil {
		log.Fatal("failed to open db", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Panic("db connected")
	return db
}
