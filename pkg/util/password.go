package util

import "golang.org/x/crypto/bcrypt"

// Encrypt 密码加密
func Encrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Decrypt 判断密码加密后的密码和加密前的密码是否相等
func Decrypt(DBPassword, passwordRequest string) error {
	return bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(passwordRequest))
}
