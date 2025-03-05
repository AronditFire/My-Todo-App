package entities

import "time"

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	IsDone      bool      `json:"done" db:"done" default:"false"`
	CreatedAt   time.Time `json:"created" db:"created"`
}

type UserTasks struct {
	ID      int
	User_ID int
	Task_ID int
}
