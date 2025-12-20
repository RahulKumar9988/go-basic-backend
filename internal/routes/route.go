package routes

import (
	"net/http"

	"github.com/RahulKumar9988/go-basic-backend/internal/handler"
)

func RegisterRoutes(auth *handler.AuthHander, todo *handler.TodoHandler) {
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/todos", todo.Create)
}
