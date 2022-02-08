/*
package domain holds user domain logic
*/
package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User aggregate root
type User struct {
	ID        string
	Email     EmailAddress
	Hash      Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

// New creates an User
func New(id string, email EmailAddress) User {
	return User{
		ID:    id,
		Email: email,
	}
}

// Register set current user email and hash+salt password
func (u *User) Register(pwd Password) error {
	// Check if email is valid
	e := u.Email.IsValid()
	if e != nil {
		return e
	}

	// Check if password is valid
	e = pwd.IsValid()
	if e != nil {
		return e
	}

	// Hash and Salt password
	pwd.HashAndSalt(bcrypt.MinCost)

	// Assign hash to user
	u.Hash = pwd

	return nil
}
