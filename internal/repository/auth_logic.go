package repository

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"gorm.io/gorm"
)

type AuthDatabase struct {
	db *gorm.DB
}

func NewAuthDataBase(db *gorm.DB) *AuthDatabase {
	return &AuthDatabase{db: db}
}

func (r *AuthDatabase) CreateUser(user entities.User) (int, error) {
	return 1, nil
}

func (r *AuthDatabase) GetLoggedUser(username, password string) entities.User {
	var user entities.User
	return user
}
