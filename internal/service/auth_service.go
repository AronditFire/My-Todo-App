package service

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/internal/repository"
	"github.com/AronditFire/TODO-APP/pkg/password"
	token "github.com/AronditFire/TODO-APP/pkg/token"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (r *AuthService) CreateUser(user entities.CreateUserRequest) (int, error) {
	hashed_pass, err := password.HashPassword(user.Password)
	if err != nil {
		return 1, err
	}
	user.Password = hashed_pass
	return r.repos.CreateUser(user)
}

func (r *AuthService) CheckUser(req entities.CreateUserRequest) (entities.User, error) {
	return r.repos.CheckUser(req)
}

func (r *AuthService) GenerateAccessToken(user entities.User) (string, error) {
	return token.GenerateAccessToken(user)
}

func (r *AuthService) GenerateRefreshToken(user entities.User) (string, error) {
	return token.GenerateRefreshToken(user)
}

func (r *AuthService) ParseTokens(AccessToken, RefreshToken string) (int, error) {
	return 1, nil
}

func (r *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	return token.RefreshAccessToken(refreshToken)
}
