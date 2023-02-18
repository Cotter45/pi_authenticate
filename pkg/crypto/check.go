package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// CheckHash compares hashed strings
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
