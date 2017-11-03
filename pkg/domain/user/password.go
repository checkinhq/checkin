package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Password represents an encrypted password.
type Password []byte

// Hash encrypts a password.
func (p *Password) Hash(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	*p = hashedPassword

	return nil
}

// Verify verifies a password.
func (p Password) Verify(password string) error {
	return bcrypt.CompareHashAndPassword(p, []byte(password))
}

// String returns the password hash in it's string representation.
func (p Password) String() string {
	return string(p)
}
