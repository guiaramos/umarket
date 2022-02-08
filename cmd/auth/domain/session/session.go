package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	auth "github.com/guiaramos/umarket/pkg/auth"
)

// Session aggregate root.
type Session struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
	auth auth.Token
}

// New returns a new Session.
func NewSession(auth auth.Token) Session {
	return Session{
		auth: auth,
	}
}

// Generate creates Access Token and assign token info to Session.
func (s *Session) Generate(userId string, ttlMinutes time.Duration) (string, error) {
	// calculate the expire time
	expTime := time.Now().Add(time.Minute * ttlMinutes)

	// Assing Expire date
	s.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expTime.Unix(),
	}

	// sign claim
	return s.auth.Sign(s)
}

// GenerateFromToken verify and assign token info to Session.
func (s *Session) GenerateFromToken(token string) error {
	// Verify and Assign to Session
	return s.auth.Verify(token, s)
}
