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

func NewTodoHandler(s *services.TodoServices) *TodoHandler {
	return &TodoHandler{Service: s}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(r.Header.Get("Authorization"))
	var t model.Todo
	json.NewDecoder(r.Body).Decode(&t)
	t.UserID = userID

	err := h.Service.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo Created"})
}
