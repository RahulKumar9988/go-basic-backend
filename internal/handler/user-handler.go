package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RahulKumar9988/go-basic-backend/internal/repository"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := repository.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func CreateUser() {

}

func UpdateUser() {

}

func GetUsers() {

}

func DeleteUser() {

}
