package user

import "fmt"

var (
	// ErrEmailInvalid is when given email is not a valid email.
	ErrEmailInvalid = fmt.Errorf("email is not a valid email")
	// ErrPasswordInvalid is when given password is not a valid password.
	ErrPasswordInvalid = fmt.Errorf("password is not a valid password")
)
