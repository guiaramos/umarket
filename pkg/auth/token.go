package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/guiaramos/umarket/pkg/apperror"
)

type Token interface {
	Sign(claims jwt.Claims) (string, *apperror.AppError)
	Verify(token string, claims jwt.Claims) *apperror.AppError
}
