/*
package user holds user domain logic
*/
package user

import (
	"time"

	"github.com/guiaramos/umarket/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

// User aggregate root
type User struct {
	ID        string       `json:"id,omitempty" bson:"_id,omitempty"`
	Email     EmailAddress `json:"email,omitempty" bson:"email"`
	Hash      Password     `json:"hash,omitempty" bson:"hash"`
	CreatedAt time.Time    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// New creates an User
func New(id string, email EmailAddress) User {
	return User{
		ID:    id,
		Email: email,
	}
}

// Register set current user email and hash+salt password
func (u *User) Register(pwd Password) *apperror.AppError {
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
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}
