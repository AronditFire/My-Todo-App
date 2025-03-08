package repository

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/pkg/password"
	"gorm.io/gorm"
)

type AuthDatabase struct {
	db *gorm.DB
}

func NewAuthDataBase(db *gorm.DB) *AuthDatabase {
	return &AuthDatabase{db: db}
}

func (r *AuthDatabase) CreateUser(user entities.CreateUserRequest) (int, error) {
	adduser := entities.User{
		Login:    user.Login,
		Password: user.Password,
	}

	var existingUser entities.User
	if err := r.db.Where("login = ?", adduser.Login).First(&existingUser).Error; err == nil { // is already in db
		return 1, err
	}

	if err := r.db.Create(&user).Error; err != nil {
		return 1, err
	}

	return int(adduser.ID), nil
}

func (r *AuthDatabase) CheckUser(req entities.CreateUserRequest) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("login = ?", user.Login).First(&user).Error; err != nil {
		return entities.User{}, err
	}

	if err := password.ComparePassword(req.Password, user.Password); err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *AuthDatabase) GetLoggedUser(username, password string) entities.User {
	var user entities.User
	return user
}
