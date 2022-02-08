package domain_test

import (
	"testing"

	user "github.com/guiaramos/umarket/cmd/user/domain"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
)

// Validate simulates password isValid call
func Validate(pwd user.Password) (err error) {
	err = pwd.IsValid()
	return err
}

func TestPassword_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		input user.Password
		err   error
	}{
		{
			name:  "should return ErrPasswordInvalid if password is less than 7 digits",
			input: "As123@",
			err:   user.ErrPasswordInvalid,
		},
		{
			name:  "should return ErrPasswordInvalid if password has no upper case char",
			input: "answer1234@",
			err:   user.ErrPasswordInvalid,
		},
		{
			name:  "should return ErrPasswordInvalid if password has no lower case char",
			input: "ANSWER1234@",
			err:   user.ErrPasswordInvalid,
		},
		{
			name:  "should return ErrPasswordInvalid if password has no number char",
			input: "AnswerHail@",
			err:   user.ErrPasswordInvalid,
		},
		{
			name:  "should return ErrPasswordInvalid if password has no special char",
			input: "Answer21051",
			err:   user.ErrPasswordInvalid,
		},
		{
			name:  "should return nil if password has passed validations",
			input: "Answer2105@",
			err:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Validate(test.input)
			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, test.err)
			}
		})
	}
}

func TestPassword_HashAndSalt(t *testing.T) {
	pass := "Answer2105@"
	var pwd user.Password = (user.Password)(pass)

	t.Run("should hash password", func(t *testing.T) {
		assert.Equal(t, pass, pwd.String())

		e := pwd.HashAndSalt(bcrypt.MinCost)
		assert.NoError(t, e)

		assert.NotEqual(t, pwd.String(), pass)

		ok := pwd.ComparePassword(pass)
		assert.True(t, ok)
	})

	t.Run("should return error if bcrypt fails", func(t *testing.T) {
		c := 32
		e := pwd.HashAndSalt(c)
		assert.ErrorIs(t, bcrypt.InvalidCostError(c), e)
	})
}

func TestPassword_ComparePassword(t *testing.T) {
	pass := "Answer2105@"
	var pwd user.Password = (user.Password)(pass)
	pwd.HashAndSalt(bcrypt.MinCost)

	t.Run("should return false if password is not matching", func(t *testing.T) {
		ok := pwd.ComparePassword("some other password")
		assert.False(t, ok)
	})

	t.Run("should return true if password is matching", func(t *testing.T) {
		ok := pwd.ComparePassword(pass)
		assert.True(t, ok)
	})
}

func TestPassword_String(t *testing.T) {
	pass := "Answer2105@"
	var pwd user.Password = (user.Password)(pass)

	t.Run("should return type string", func(t *testing.T) {
		p := pwd.String()
		assert.Equal(t, pass, p)
	})
}
