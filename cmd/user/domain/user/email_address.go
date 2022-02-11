package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/guiaramos/umarket/pkg/apperror"
)

// EmailAddress is an email address.
type EmailAddress string

// IsValid returns error if string is not valid email.
func (e EmailAddress) IsValid() *apperror.AppError {
	if !govalidator.IsEmail(string(e)) {
		return ErrEmailInvalid
	}

	return nil
}
