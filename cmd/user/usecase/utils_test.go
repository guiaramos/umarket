package usecase_test

import (
	"context"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/pkg/apperror"
)

type SpyUserRepository struct {
	insertOneWithError   bool
	updateOneWithError   bool
	findOneWithError     bool
	findByEmailWithError bool

	findByEmailFound bool

	newIdWasCalled       bool
	insertOneWasCalled   bool
	updateOneWasCalled   bool
	findOneWasCalled     bool
	findByEmailWasCalled bool
}

const (
	email    user.EmailAddress = "gui_aramos@outlook.com"
	password user.Password     = "Ash2109@"
	id                         = "507f191e810c19729de860ea"
)

func (r *SpyUserRepository) NewId() string {
	r.newIdWasCalled = true
	return id
}
func (r *SpyUserRepository) InsertOne(ctx context.Context, u *user.User) *apperror.AppError {
	r.insertOneWasCalled = true
	if r.insertOneWithError {
		return mockError()
	}
	return nil
}
func (r *SpyUserRepository) UpdateOne(ctx context.Context, u *user.User) *apperror.AppError {
	r.updateOneWasCalled = true
	if r.updateOneWithError {
		return mockError()
	}
	return nil
}
func (r *SpyUserRepository) FindOne(ctx context.Context, id string) (*user.User, *apperror.AppError) {
	r.findOneWasCalled = true
	if r.findOneWithError {
		return nil, mockError()
	}

	u := &user.User{
		ID:    id,
		Email: email,
		Hash:  password,
	}
	return u, nil
}
func (r *SpyUserRepository) FindByEmail(ctx context.Context, email user.EmailAddress) (*user.User, *apperror.AppError) {
	r.findByEmailWasCalled = true
	if r.findByEmailWithError {
		return nil, mockError()
	}

	u := &user.User{}

	if r.findByEmailFound {
		u.ID = id
		u.Email = email
		u.Hash = password
		return u, nil
	}

	return nil, nil
}

func mockError() *apperror.AppError {
	return apperror.NewInternalServerError("something went wrong")
}
