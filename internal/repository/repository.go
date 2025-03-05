package repository

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetLoggedUser(username, password string) entities.User
}

type Task interface {
	CreateTask(userID int, task entities.Task) (int, error)
	GetAllTasks(userID int) ([]entities.Task, error)
	UpdateTask(userID, taskID int, updateBody entities.UpdateTask) (int, error)
	DeleteTask(userID, taskID int) error
}

type JSON_File interface {
	AddFile() error
}

type Repository struct {
	Authorization
	Task
	JSON_File
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDataBase(db),
		Task:          NewTaskDataBase(db),
		JSON_File:     NewJSONDatabase(db),
	}
}
