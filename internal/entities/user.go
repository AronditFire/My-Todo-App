package entities

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" default:"user"`
}
