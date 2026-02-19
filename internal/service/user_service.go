package service //handles business logic

import (
	"fmt"
	"github.com/mithun/gin/internal/repository"
) //service depend on these model because it needs to use them to implement business logic

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}

	user := s.repo.Create(name)
	return user.Name, nil
}
