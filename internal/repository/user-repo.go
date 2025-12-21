// package definition
package repository

import (
	"database/sql"

	"github.com/RahulKumar9988/go-basic-backend/internal/model"
)

// connect repo -> DB
type UserRepo struct {
	DB *sql.DB
}

// constructor fucntion creation -> create and return urse instances
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// main logic
func (r *UserRepo) Create(u model.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (email, password) VALUES ($1,$2)",
		u.Email, u.Password,
	)
	return err
}

func (r *UserRepo) GetAllUsers() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, email FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Email)
		users = append(users, u)
	}
	return users, nil
}
