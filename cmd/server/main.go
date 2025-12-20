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

	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	println("env loaded")

	//Load Application
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to Load Config", err)
	}
	println("config loaded")

	// Conntect DB
	dbConnection, err := db.ConntectDB(cfg)
	if err != nil {
		log.Fatal("Failed to Connect DB:", err)
	}
	defer dbConnection.Close()

	//load Repos
	userRepo := repository.NewUserRepo(dbConnection)
	todoRepo := repository.NewTodoRepo(dbConnection)

	//load Services
	authServices := services.NewAuthService(userRepo)
	todoServices := services.NewTodoServices(todoRepo)

	// load handlers
	authHandler := handler.NewAuthHandler(authServices)
	todoHandler := handler.NewTodoHandler(todoServices)

	//connect Routes
	routes.RegisterRoutes(authHandler, todoHandler)

	// SetUP server with PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server running on port : ", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
