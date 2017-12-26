package domain

import (
	"time"

	"github.com/segmentio/ksuid"
)

// User is the central class in the domain model.
type User struct {
	// ID is the internal identifier
	ID int64

	// UID is the public unique identifier of the user
	UID ksuid.KSUID

	Email     string
	Password  Password
	FirstName string
	LastName  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser returns a new user.
func NewUser() *User {
	return &User{
		UID: ksuid.New(),
	}
}

// UserRepository provides access to a user store.
type UserRepository interface {
	// FindByID finds a user based on it's internal ID.
	FindByID(id int64) (*User, error)

	// FindByUID finds a user based on it's public unique ID.
	FindByUID(uid ksuid.KSUID) (*User, error)

	// FindByEmail finds a user based on it's email address.
	FindByEmail(email string) (*User, error)

	// Store creates or updates a user.
	Store(user *User) error
}
