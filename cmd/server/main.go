package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
	"github.com/RahulKumar9988/go-basic-backend/internal/db"
	"github.com/RahulKumar9988/go-basic-backend/internal/handler"
	"github.com/RahulKumar9988/go-basic-backend/internal/repository"
	"github.com/RahulKumar9988/go-basic-backend/internal/routes"
	"github.com/RahulKumar9988/go-basic-backend/internal/services"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading: ENV file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	dbConn, err := db.ConntectDB(cfg)
	if err != nil {
		log.Fatal("Failed ton conntect db", err)
	}

	defer dbConn.Close()
	log.Println("db connected")

	userRepo := repository.NewUserRepo(dbConn)
	todoRepo := repository.NewTodoRepo(dbConn)

	userService := services.NewUserServices(userRepo)
	todoService := services.NewTodoServices(todoRepo)

	userHandler := handler.NewUserHandler(userService)
	todoHandler := handler.NewTodoHandler(todoService)

	routes.RegisterRoutes(userHandler, todoHandler)

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("PORT missing")
		port = "8080"
	}

	log.Printf("Server is running on port: %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
