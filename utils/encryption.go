package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	securePass, parseErr := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(securePass), parseErr
}

func ComparePasswords(encryptedPassword, password string) bool {
	passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return passwordMatchErr == nil
}
