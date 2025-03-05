package service

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/internal/repository"
)

type AuthService struct {
	repos *repository.Authorization
}

func NewAuthService(repos *repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (r *AuthService) CreateUser(user entities.User) (int, error) {
	return 1, nil
}
