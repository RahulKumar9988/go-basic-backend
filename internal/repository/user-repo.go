package repository

import (
	"github.com/RahulKumar9988/go-basic-backend/internal/db"
	"github.com/RahulKumar9988/go-basic-backend/internal/model"
)

func GetAllUsers() ([]model.User, error) {
	rows, error := db.DB.Query("SELECT id, name, email from users")
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var u model.User
		rows.Scan(&u.Id, &u.Email, &u.Name)
		users = append(users, u)
	}
	return users, nil
}

func CreateUser(user *model.User) error {
	result, error := db.DB.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		user.Name,
		user.Email,
	)

	if error != nil {
		return error
	}

	id, _ := result.LastInsertId()
	user.Id = int(id)

	return nil
}

func GetUser(id int) (model.User, error) {
	var user model.User

	error := db.DB.QueryRow(
		"SELECT id, name, email FROM users where id=?",
		id,
	).Scan(&user.Id, &user.Email, &user.Name)

	return user, error
}

func UpdateUser(id int, user model.User) error {
	_, error := db.DB.Exec(
		"UPDATE users SET name=?, email=? WHERE id=?",
		user.Email,
		user.Name,
		id,
	)
	return error
}

func DeleteUser(id int) error {
	_, error := db.DB.Exec(
		"DELETE FROM users WHERE id?",
		id,
	)

	return error
}
