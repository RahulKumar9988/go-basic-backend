package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
)

var DB *sql.DB

func ConntectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBHost,
		cfg.DBPass,
		cfg.DBUser,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
