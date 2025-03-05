package repository

import (
	"gorm.io/gorm"
)

type JSONDatabase struct {
	db *gorm.DB
}

func NewJSONDatabase(db *gorm.DB) *JSONDatabase {
	return &JSONDatabase{db: db}
}

func (j *JSONDatabase) AddFile() error {
	return nil
}
