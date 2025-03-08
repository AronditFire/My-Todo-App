package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:varchar(50);not null;default:'user'"`
}

type CreateUserRequest struct {
	Login    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"` // min 6 symbols
}
