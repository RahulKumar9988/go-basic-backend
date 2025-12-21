package services

import "github.com/RahulKumar9988/go-basic-backend/internal/repository"

type UserServices struct {
	Repo *repository.UserRepo
}

func NewUserServices(r *repository.UserRepo) *UserServices {
	return &UserServices{Repo: r}
}

func (s *UserServices) 
