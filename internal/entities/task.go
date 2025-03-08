package entities

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255)"`
	IsDone      bool   `gorm:"type:boolean;default:false"`
}

type UserTasks struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	ListID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"` // TODO: CHECK
	List   Task `gorm:"foreignKey:ListID;references:ID;constraint:OnDelete:CASCADE"`
}

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}
