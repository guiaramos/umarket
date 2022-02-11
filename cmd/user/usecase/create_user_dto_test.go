package usecase_test

import (
	"testing"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/cmd/user/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserDTO_NewCreateUserDTO(t *testing.T) {
	t.Run("should create a new CreateUserDTO", func(t *testing.T) {
		dto := usecase.NewCreateUserDTO(email, password)

		assert.Equal(t, email, dto.Email)
		assert.Equal(t, password, dto.Password)
	})
}

func TestCreateUserDTO_Validate(t *testing.T) {
	t.Run("should return nil if all data is valid", func(t *testing.T) {
		dto := usecase.NewCreateUserDTO(email, password)

		err := dto.Validate()

		assert.Nil(t, err)
	})

	t.Run("should return error if email is not valid", func(t *testing.T) {
		var wrongEmail user.EmailAddress = "not a email"
		dto := usecase.NewCreateUserDTO(wrongEmail, password)

		err := dto.Validate()

		assert.ErrorIs(t, user.ErrEmailInvalid, err)
	})

	t.Run("should return error if password is not valid", func(t *testing.T) {
		var wrongPassword user.Password = "weak password"
		dto := usecase.NewCreateUserDTO(email, wrongPassword)

		err := dto.Validate()

		assert.ErrorIs(t, user.ErrPasswordInvalid, err)
	})
}
