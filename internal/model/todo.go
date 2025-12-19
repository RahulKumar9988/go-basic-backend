package model

type Todo struct {
	Id     string `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}
