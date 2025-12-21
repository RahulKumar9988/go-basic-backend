package routes

import (
	"net/http"

	"github.com/RahulKumar9988/go-basic-backend/internal/handler"
)

func RegisterRoutes(
	user *handler.UserHandler,
	todo *handler.TodoHandler,
) {
	http.HandleFunc("/users", user.Create)
	http.HandleFunc("/todos", todo.Create)
}
