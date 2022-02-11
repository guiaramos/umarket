package user

import (
	"unicode"

	"github.com/guiaramos/umarket/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

// Password is an email address.
type Password string

// IsValid returns error if string is not valid password.
func (p Password) IsValid() *apperror.AppError {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(p) >= 7 {
		hasMinLen = true
	}
	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
		return nil
	}

	return ErrPasswordInvalid
}

// HashAndSalt hash and salt the password
func (p *Password) HashAndSalt(cost int) *apperror.AppError {
	bytePwd := []byte(p.String())
	hash, err := bcrypt.GenerateFromPassword(bytePwd, cost)
	if err != nil {
		return apperror.NewInternalServerError(err.Error())
	}

	*p = (Password)(hash)

	return nil
}

// ComparePassword verifies if password matched with hash
func (h Password) ComparePassword(pwd string) bool {
	bytePwd := []byte(pwd)
	byteHash := []byte(h)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)

	return err == nil
}

func (p Password) String() string {
	return string(p)
}
