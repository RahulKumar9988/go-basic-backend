package services

import (
	"github.com/RahulKumar9988/go-basic-backend/internal/model"
	"github.com/RahulKumar9988/go-basic-backend/internal/repository"
)

type UserServices struct {
	Repo *repository.UserRepo
}

func NewUserServices(r *repository.UserRepo) *UserServices {
	return &UserServices{Repo: r}
}

func (s *UserServices) CreateUser(u model.User) error {
	return s.Repo.Create(u)
}

func (s *UserServices) GetUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}
