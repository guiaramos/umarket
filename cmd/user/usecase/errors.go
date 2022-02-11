package usecase

import (
	"github.com/guiaramos/umarket/pkg/apperror"
)

var (
	// ErrEmailAlreadyInUse is when given email is already in use.
	ErrEmailAlreadyInUse = apperror.NewBadRequest("email address is already in use")
)
