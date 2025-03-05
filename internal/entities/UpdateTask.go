package entities

import "errors"

type UpdateTask struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsDone      *bool   `json:"done"`
}

func (u *UpdateTask) IsEmptyUpdate() error {
	if u.Description == nil && u.Title == nil && u.IsDone == nil {
		return errors.New("updated fields not found")
	}

	return nil
}
