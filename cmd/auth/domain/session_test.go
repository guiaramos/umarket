package auth_test

import (
	"testing"

	"github.com/golang-jwt/jwt"
	auth "github.com/guiaramos/umarket/cmd/auth/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userId = "some_user_id"
	token  = "fake_token"
)

type SpyAuthToken struct {
	signWasCalled   bool
	verifyWasCalled bool
}

func (s *SpyAuthToken) Sign(claims jwt.Claims) (string, error) {
	s.signWasCalled = true
	return token, nil
}
func (s *SpyAuthToken) Verify(token string, claims jwt.Claims) error {
	s.verifyWasCalled = true
	return nil
}

func TestSession_NewSession(t *testing.T) {
	authSpy := &SpyAuthToken{}

	t.Run("should create without UserId", func(t *testing.T) {
		session := auth.NewSession(authSpy)

		assert.Empty(t, session.UserID)
	})
}

func TestSession_Generate(t *testing.T) {
	authSpy := &SpyAuthToken{}
	session := auth.NewSession(authSpy)

	t.Run("should assign user and expire to session", func(t *testing.T) {
		accessToken, err := session.Generate(userId, 5)
		assert.NoError(t, err)
		assert.Equal(t, accessToken, token)
	})

	t.Run("should generate token without error", func(t *testing.T) {
		accessToken, err := session.Generate(userId, 5)

		assert.True(t, authSpy.signWasCalled)
		assert.NoError(t, err)
		assert.Equal(t, accessToken, token)
	})

}

func TestSession_GenerateFromToken(t *testing.T) {
	authSpy := &SpyAuthToken{}
	session := auth.NewSession(authSpy)
	session.Generate(userId, 5)

	t.Run("should parse token and assign value", func(t *testing.T) {
		session.GenerateFromToken(token)
		assert.True(t, authSpy.verifyWasCalled)
	})
}
