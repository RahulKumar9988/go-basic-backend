package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RahulKumar9988/go-basic-backend/internal/model"
	"github.com/RahulKumar9988/go-basic-backend/internal/services"
)

type TodoHandler struct {
	Service *services.TodoServices
}

func NewTodoHandler(services *services.TodoServices) *TodoHandler {
	return &TodoHandler{Service: services}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("Autharization")
	if userIDStr != "" {
		http.Error(w, "unautharization", http.StatusUnauthorized)
		return
	}

	userID, _ := strconv.Atoi(userIDStr)

	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.UserId = userID

	h.Service.Create(todo)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Todo Created",
	})
}
