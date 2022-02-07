package auth

import "fmt"

var (
	// ErrTokenInvalid is when given token is not valid.
	ErrTokenInvalid = fmt.Errorf("token is not valid")
)
