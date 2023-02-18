package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// CreateHash creates a hashed string
func CreateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
