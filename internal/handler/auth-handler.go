// package defination
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RahulKumar9988/go-basic-backend/internal/services"
)

// create conection between handler and service
type AuthHander struct {
	Service *services.AuthService
}

// create constructore for create safe object
func NewAuthHandler(services *services.AuthService) *AuthHander {
	return &AuthHander{Service: services}
}

// main fucntion logic
func (h *AuthHander) Login(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	json.NewDecoder(r.Body).Decode(&body)

	userID, err := h.Service.Login(body["email"], body["password"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// simple token = userID (for learning)
	json.NewEncoder(w).Encode(map[string]string{
		"token": strconv.Itoa(userID),
	})
}
