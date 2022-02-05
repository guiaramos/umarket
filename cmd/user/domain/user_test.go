package user_test

import (
	"testing"

	user "github.com/guiaramos/umarket/cmd/user/domain"
	apptest "github.com/guiaramos/umarket/pkg/testing"
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
		apptest.AssertError(t, e, user.ErrEmailInvalid)
	})

	t.Run("should return error if password is not valid", func(t *testing.T) {
		u := user.New(id, email)

		e := u.Register("not good password")
		apptest.AssertError(t, e, user.ErrPasswordInvalid)
	})

	t.Run("should create a new user", func(t *testing.T) {
		u := user.New(id, email)

		e := u.Register(pwd)
		apptest.AssertNoError(t, e)

		if u.ID == "" {
			t.Fatal("expected to have an id but didn't get one")
		}

		if u.Email == "" {
			t.Fatal("expected to have an email but didn't get one")
		}

		if u.Hash == "" {
			t.Fatal("expected to have a hash but didn't get one")
		}

		if u.Hash == pwd {
			t.Fatal("expected hash not to be equal to the password")
		}

		if ok := u.Hash.ComparePassword(pwd.String()); !ok {
			t.Fatal("expected hash to match password")
		}
	})
}
