// services/userService.go
package services

import (
	"go-backend-portfolio/models"
	"go-backend-portfolio/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.Repo.CreateUser(user)
}
