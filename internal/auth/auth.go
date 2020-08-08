package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password and returns the byte value
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 8)
}

// Authenticate compares two passwords and returns result
func Authenticate(hashedPassword, enteredPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}
