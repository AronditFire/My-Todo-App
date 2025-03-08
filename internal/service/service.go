package service

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/internal/repository"
)

type Authorization interface {
	CreateUser(user entities.CreateUserRequest) (int, error)
	CheckUser(req entities.CreateUserRequest) (entities.User, error)
	GenerateAccessToken(user entities.User) (string, error)
	GenerateRefreshToken(user entities.User) (string, error)
	ParseTokens(AccessToken, RefreshToken string) (int, error)
	RefreshAccessToken(refreshToken string) (string, error)
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

type Service struct {
	Authorization
	Task
	JSON_File
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Task:          NewTaskService(repos.Task),
		JSON_File:     NewJSONService(repos.JSON_File),
	}
}
