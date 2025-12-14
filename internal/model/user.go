package model

// user model
type User struct {
	ID        string `json:"id" db:"id"`
	USserName string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
}

// todos model
type Todos struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status bool   `json:"status" db:"status"`
}
