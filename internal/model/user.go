package model

// user model
type User struct {
	ID        string `json:"id" db:"id"`
	USserName string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
}
