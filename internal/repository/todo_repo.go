package repository

import (
	"database/sql"

	"github.com/RahulKumar9988/go-basic-backend/internal/model"
)

type TodoRepo struct {
	DB *sql.DB
}

func NewTodoRepo(db *sql.DB) *TodoRepo {
	return &TodoRepo{DB: db}
}

// select by ID
func (r TodoRepo) GetByUser(userID int) ([]model.Todo, error) {
	rows, err := r.DB.Query(
		"SELECT id, title, completed ,user_id FROM todos WHERE user_id=$1",
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		rows.Scan(&t.Id, &t.Title, &t.Completed, &t.UserID)
		todos = append(todos, t)
	}

	return todos, nil
}

// create operation
func (r *TodoRepo) Create(todo model.Todo) error {
	_, err := r.DB.Exec(
		"INSERT INTO todos (title, completed, user_id) VALUES ($1, $2, $3)",
		todo.Title,
		todo.Completed,
		todo.UserID,
	)
	return err
}

// update operation
func (r *TodoRepo) Update(id string, todo model.Todo) error {
	_, err := r.DB.Exec(
		"UPDATE todos SET title=?, status=? WHERE id=?",
		todo.Title,
		// todo.Status,
		id,
	)

	return err
}

// delete operation
func (r TodoRepo) Delete(id string) error {
	_, err := r.DB.Exec(
		"DELETE FROM todos WHERE id=?",
		id,
	)

	return err
}
