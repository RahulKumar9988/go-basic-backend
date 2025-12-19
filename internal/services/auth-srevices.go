// package defined
package services

import (
	"errors"

	"github.com/RahulKumar9988/go-basic-backend/internal/repository"
)

// service creation
type AuthService struct {
	Repo *repository.UserRepo
}

// inject Dependency -> connect repo
func NewAuthService(repo *repository.UserRepo) *AuthService {
	return &AuthService{Repo: repo}
}

// acess login
func (s *AuthService) Login(email, password string) (string, error) {

	// fetch user from db
	user, err := s.Repo.FindByEmail(email)

	// check user exist of not
	if err != nil {
		return "", errors.New("User Not Found")
	}

	// verify password
	if user.Password != password {
		return "", errors.New("Invalid Password")
	}

	// return login user id
	return user.ID, nil
}
