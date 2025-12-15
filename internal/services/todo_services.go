package services

import (
	"github.com/RahulKumar9988/go-basic-backend/internal/model"
	"github.com/RahulKumar9988/go-basic-backend/internal/repository"
)

type TodoServices struct {
	Repo *repository.TodoRepo
}

func NewTodoServices(repo *repository.TodoRepo) *TodoServices {
	return &TodoServices{Repo: repo}
}

func (s *TodoServices) Create(todo model.Todos) error {
	return s.Repo.Create(todo)
}

func (s *TodoServices) GetAll() ([]model.Todos, error) {
	return s.Repo.GetAll()
}

func (s *TodoServices) Update(id string, todo model.Todos) error {
	return s.Repo.Update(id, todo)
}

func (s *TodoServices) Delete(id string) error {
	return s.Repo.Delete(id)
}
