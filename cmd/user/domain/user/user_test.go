package user_test

import (
	"testing"

	user "github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/teris-io/shortid"
)

var (
	id                      = shortid.MustGenerate()
	email user.EmailAddress = "gui_aramos@outlook.com"
	pwd   user.Password     = "Answer2105@"
)

func TestUser_Register(t *testing.T) {

	t.Run("should return error if email is not valid", func(t *testing.T) {
		u := user.New(id, "no good email")

		e := u.Register(pwd)
		assert.ErrorIs(t, user.ErrEmailInvalid, e)
	})

	t.Run("should return error if password is not valid", func(t *testing.T) {
		u := user.New(id, email)

		e := u.Register("not good password")
		assert.ErrorIs(t, user.ErrPasswordInvalid, e)
	})

	t.Run("should create a new user", func(t *testing.T) {
		u := user.New(id, email)

		e := u.Register(pwd)
		assert.NoError(t, e)
		assert.NotEmpty(t, u.ID)
		assert.NotEmpty(t, u.Email)
		assert.NotEmpty(t, u.Hash)
		assert.NotEqual(t, pwd, u.Hash)
		assert.True(t, u.Hash.ComparePassword(pwd.String()))
	})
}
