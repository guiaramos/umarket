package auth

import (
	"github.com/golang-jwt/jwt"
)

type Token interface {
	Sign(claims jwt.Claims) (string, error)
	Verify(token string, claims jwt.Claims) error
}
