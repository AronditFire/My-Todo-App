package service

import "github.com/AronditFire/TODO-APP/internal/repository"

type JSONService struct {
	repos repository.JSON_File
}

func NewJSONService(repos repository.JSON_File) *JSONService {
	return &JSONService{repos: repos}
}

func (r *JSONService) AddFile() error {
	return nil
}
