package user

import (
	"github.com/guiaramos/umarket/pkg/apperror"
)

var (
	// ErrEmailInvalid is when given email is not a valid email.
	ErrEmailInvalid = apperror.NewBadRequest("email is not a valid email")
	// ErrPasswordInvalid is when given password is not a valid password.
	ErrPasswordInvalid = apperror.NewBadRequest("password is not a valid password")
)
