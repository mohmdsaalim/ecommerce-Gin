package utils

import "golang.org/x/crypto/bcrypt"

// password hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
// checking password
func Checkpassword(hash, password string)bool  {
	return  bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))==nil
}