package user

import (
	"context"

	"github.com/guiaramos/umarket/pkg/apperror"
)

// Repository allows persistent operations for User
type Repository interface {
	NewId() string
	InsertOne(ctx context.Context, u *User) *apperror.AppError
	UpdateOne(ctx context.Context, u *User) *apperror.AppError
	FindOne(ctx context.Context, id string) (*User, *apperror.AppError)
	FindByEmail(ctx context.Context, email EmailAddress) (*User, *apperror.AppError)
}
