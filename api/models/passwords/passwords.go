package passwords

import "golang.org/x/crypto/bcrypt"

const (
	salt = 10
)

// Encrypt - takes a password and  encrypt it
func Encrypt(pass string) (string, error) {
	val, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(val), err
}

// Compare - this compares the hashed value in DB to the passed in value of the endpoint
func Compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
