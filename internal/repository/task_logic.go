package repository

import (
	"github.com/AronditFire/TODO-APP/internal/entities"
	"gorm.io/gorm"
)

type TaskDatabase struct {
	db *gorm.DB
}

func NewTaskDataBase(db *gorm.DB) *TaskDatabase {
	return &TaskDatabase{db: db}
}

func (t *TaskDatabase) CreateTask(userID int, task entities.Task) (int, error) {
	return 1, nil
}

func (t *TaskDatabase) GetAllTasks(userID int) ([]entities.Task, error) {
	var items []entities.Task
	return items, nil
}

func (t *TaskDatabase) UpdateTask(userID, taskID int, updateBody entities.UpdateTask) (int, error) {
	return 1, nil
}

func (t *TaskDatabase) DeleteTask(userID, taskID int) error {
	return nil
}
