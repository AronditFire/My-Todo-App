package service

import "github.com/AronditFire/TODO-APP/internal/entities"

type Authorization interface {
	CreateUser(user entities.User) (int, error)
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
