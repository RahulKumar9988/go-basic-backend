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
func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	// query is required to get the user by id.
	row := r.DB.QueryRow(
		"SELECT id, email, password FROM users WHERE email=?",
		email,
	)

	// empty struct created to store user
	var user model.User

	// scane db and map it with user struct
	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	// If no row found OR any DB error occurs
	if err != nil {
		return nil, err
	}

	// return pointer to user struct if sucessful
	return &user, nil
}
