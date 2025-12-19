package main

import (
	"fmt"
	"net/http"

	"github.com/RahulKumar9988/go-basic-backend/internal/config"
)

func main() {
	const PORT int = 3000
	fmt.Println("Your server running on port", PORT)

	//load config
	cfg := config.LoadConfig()
	fmt.Println("DB Host is:", cfg.DBHost)

	// server setup
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, to todo manager!"))
	})

}
