package model

type Todo struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserId string `json:"user_id"`
}
