package auth

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/guiaramos/umarket/pkg/apperror"
)

const (
	privKeyPath = "cert/id.rsa"     // openssl genrsa -out id.rsa -2048
	pubKeyPath  = "cert/id.rsa.pub" // openssl rsa -in id.rsa -pubout > id.rsa.pub
)

var (
	signKey   []byte
	verifyKey []byte
)

type jwtToken struct {
}

func init() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	fmt.Println(err)
	fatal(err)

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	fatal(err)
}

// Sign generates token string
func (j jwtToken) Sign(claims jwt.Claims) (string, *apperror.AppError) {
	// create jwt with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return the signed string
	signed, err := token.SignedString(signKey)
	if err != nil {
		return "", apperror.NewInternalServerError(err.Error())
	}

	return signed, nil
}

// Verify checks if token is valid and assign the claims
func (j jwtToken) Verify(token string, claims jwt.Claims) *apperror.AppError {
	decoded, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		return apperror.NewUnauthorized(err.Error())
	}

	if !decoded.Valid {
		return ErrTokenInvalid
	}

	return nil
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
