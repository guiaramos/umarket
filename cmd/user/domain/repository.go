package domain

import "context"

// Repository allows persistent operations for User
type Repository interface {
	Save(ctx context.Context, u *User) error
	Get(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
