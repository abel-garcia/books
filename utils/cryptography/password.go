package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = "_kYOD3igoAaiBw=="
)

func CreatePasswordWithSalt(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(hashPassword), err
}

func ComparePasswordWithSalt(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
}
