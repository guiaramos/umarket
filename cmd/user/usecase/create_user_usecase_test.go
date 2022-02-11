package usecase_test

import (
	"context"
	"testing"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/cmd/user/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_NewCreateUserUseCase(t *testing.T) {
	t.Run("should create new CreateUserUseCase", func(t *testing.T) {
		repo := &SpyUserRepository{}
		uc := usecase.NewCreateUserUseCase(repo)

		assert.NotNil(t, uc)
	})
}

func TestCreateUserUseCase_Execute(t *testing.T) {
	repo := &SpyUserRepository{}
	uc := usecase.NewCreateUserUseCase(repo)
	dto := usecase.NewCreateUserDTO(email, password)

	t.Run("should return error if find by email fails", func(t *testing.T) {
		r := &SpyUserRepository{
			findByEmailWithError: true,
		}
		c := usecase.NewCreateUserUseCase(r)
		err := c.Execute(context.TODO(), dto)

		assert.True(t, r.findByEmailWasCalled)
		assert.NotNil(t, err)
		assert.False(t, r.newIdWasCalled)
		assert.False(t, r.insertOneWasCalled)
	})

	t.Run("should check if email is already in use", func(t *testing.T) {
		r := &SpyUserRepository{
			findByEmailFound: true,
		}
		c := usecase.NewCreateUserUseCase(r)
		err := c.Execute(context.TODO(), dto)

		assert.True(t, r.findByEmailWasCalled)
		assert.Equal(t, usecase.ErrEmailAlreadyInUse, err)
		assert.False(t, r.newIdWasCalled)
		assert.False(t, r.insertOneWasCalled)
	})

	t.Run("should return error if register fails", func(t *testing.T) {
		dto := usecase.NewCreateUserDTO(email, "wrong password")
		err := uc.Execute(context.TODO(), dto)

		assert.True(t, repo.findByEmailWasCalled)
		assert.True(t, repo.newIdWasCalled)
		assert.Equal(t, user.ErrPasswordInvalid, err)
		assert.False(t, repo.insertOneWasCalled)
	})

	t.Run("should return error if InsertOne fails", func(t *testing.T) {
		r := &SpyUserRepository{
			insertOneWithError: true,
		}
		c := usecase.NewCreateUserUseCase(r)
		err := c.Execute(context.TODO(), dto)

		assert.True(t, r.findByEmailWasCalled)
		assert.True(t, r.newIdWasCalled)
		assert.True(t, r.insertOneWasCalled)
		assert.NotNil(t, err)
	})

	t.Run("should create new user without errors", func(t *testing.T) {
		err := uc.Execute(context.TODO(), dto)

		assert.Nil(t, err)
	})
}
