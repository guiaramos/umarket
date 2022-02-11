package user_test

import (
	"testing"

	user "github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestEmailAddress_IsValid(t *testing.T) {
	t.Run("should return ErrEmailInvalid if email is not valid", func(t *testing.T) {
		var email user.EmailAddress = "gui_aramos"
		err := email.IsValid()

		assert.ErrorIs(t, err, user.ErrEmailInvalid)
	})

	t.Run("should return nil if email is valid", func(t *testing.T) {
		var email user.EmailAddress = "gui_aramos@outlook.com"
		err := email.IsValid()

		assert.Nil(t, err)
	})

}
