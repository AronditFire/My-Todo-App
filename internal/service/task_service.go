package service

import "github.com/AronditFire/TODO-APP/internal/repository"

type TaskService struct {
	repos repository.Task
}

func NewTaskService(repos repository.Task) *TaskService {
	return &TaskService{repos: repos}
}
