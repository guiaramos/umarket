package user

import "context"

// Repository allows persistent operations for User
type Repository interface {
	NewId() string
	InsertOne(ctx context.Context, u *User) error
	UpdateOne(ctx context.Context, u *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email EmailAddress) (*User, error)
}
