package service

import (
	"userOnboard/models"
	"userOnboard/repository"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUserByID(id string) (*models.User, error)
	ListUsers() ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}
