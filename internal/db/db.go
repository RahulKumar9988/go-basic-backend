package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib" 
)

func ConntectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=require",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
