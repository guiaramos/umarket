package auth

import (
	"github.com/guiaramos/umarket/pkg/apperror"
)

var (
	// ErrTokenInvalid is when given token is not valid.
	ErrTokenInvalid = apperror.NewUnauthorized("token is not valid")
)
