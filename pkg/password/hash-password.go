package password

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(password), 10) // salt length 10
	return string(hashedpass), err
}

func ComparePassword(password string, hashedpass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
	return err
}
