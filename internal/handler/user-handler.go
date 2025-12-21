package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RahulKumar9988/go-basic-backend/internal/model"
	"github.com/RahulKumar9988/go-basic-backend/internal/services"
)

type UserHandler struct {
	Service *services.UserServices
}

func NewUserHandler(s *services.UserServices) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User
	json.NewDecoder(r.Body).Decode(&u)

	h.Service.CreateUser(u)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Created"})
}
