package service

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/internal/repository"
)

type TaskService struct {
	repos repository.Task
}

func NewTaskService(repos repository.Task) *TaskService {
	return &TaskService{repos: repos}
}

func (r *TaskService) CreateTask(userID int, task entities.Task) (int, error) {
	return 1, nil
}

func (r *TaskService) GetAllTasks(userID int) ([]entities.Task, error) {
	var task []entities.Task
	return task, nil
}

func (r *TaskService) UpdateTask(userID, taskID int, updateBody entities.UpdateTask) (int, error) {
	return 1, nil
}

func (r *TaskService) DeleteTask(userID, taskID int) error {
	return nil
}
