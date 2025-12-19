package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
)

var DB *sql.DB

func DbConnection(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBHost,
		cfg.DBPass,
		cfg.DBUser,
		cfg.DBPort,
		cfg.DBName,
	)

	var error error

	DB, error = sql.Open("mySql", dsn)

	if error != nil {
		log.Fatal("error connecting db connection:", error)
	}

	if error = DB.Ping(); error != nil {
		log.Fatal("failed to connect to mySql:", error)
	}

	log.Println("mySql connected sucessfully")
}
