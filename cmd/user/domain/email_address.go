package user

import "github.com/asaskevich/govalidator"

// EmailAddress is an email address.
type EmailAddress string

// IsValid returns error if string is not valid email.
func (e EmailAddress) IsValid() error {
	if !govalidator.IsEmail(string(e)) {
		return ErrEmailInvalid
	}

	return nil
}
