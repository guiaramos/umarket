package usecase

import (
	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/pkg/apperror"
)

// CreateUserDTO defines data required for creating an user.
type CreateUserDTO struct {
	Email    user.EmailAddress `json:"email,omitempty"`
	Password user.Password     `json:"password,omitempty"`
}

// NewCreateUserDTO creates a CreateUserDTO.
func NewCreateUserDTO(email user.EmailAddress, password user.Password) *CreateUserDTO {
	return &CreateUserDTO{
		Email:    email,
		Password: password,
	}
}

// Validate checks if the DTO has valid data
func (c CreateUserDTO) Validate() (err *apperror.AppError) {
	if err = c.Email.IsValid(); err != nil {
		return err
	}

	if err = c.Password.IsValid(); err != nil {
		return err
	}

	return nil
}
