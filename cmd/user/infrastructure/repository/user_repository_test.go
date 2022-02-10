package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/cmd/user/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_NewUserRepository(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	mt.Run("should return user repository", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		assert.NotEmpty(t, repo)
	})
}

func TestUserRepository_NewId(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	mt.Run("should return a new id", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)
		id := repo.NewId()

		assert.True(t, primitive.IsValidObjectID(id))
	})
}

func TestUserRepository_InsertOne(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	u := &user.User{
		ID:        primitive.NewObjectID().Hex(),
		Email:     "gui_aramos@outlook.com",
		Hash:      "mock_hash",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mt.Run("should create with success", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.InsertOne(context.TODO(), u)

		assert.NoError(t, err)
	})

	mt.Run("should return error on mongodb error", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := repo.InsertOne(context.TODO(), u)

		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestUserRepository_UpdateOne(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	u := &user.User{
		ID:        primitive.NewObjectID().Hex(),
		Email:     "gui_aramos@outlook.com",
		Hash:      "mock_hash",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mt.Run("should update one with success", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: u.ID},
				{Key: "email", Value: u.Email},
			}},
		})

		err := repo.UpdateOne(context.TODO(), u)

		assert.NoError(t, err)
	})

	mt.Run("should return error if id is not valid ObjectID", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    "wrong id",
			Email: "gui_aramos@outlook.com",
		}

		err := repo.UpdateOne(context.TODO(), u)

		assert.ErrorIs(t, primitive.ErrInvalidHex, err)
	})

	mt.Run("should return error on mongodb error", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := repo.UpdateOne(context.TODO(), u)

		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestUserRepository_FindOne(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	mt.Run("should find one with success", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: u.ID},
			{Key: "email", Value: u.Email},
		}))

		user, err := repo.FindOne(context.TODO(), u.ID)

		assert.NoError(t, err)
		assert.Equal(t, u, user)
	})

	mt.Run("should return error if id is not valid ObjectID", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    "wrong id",
			Email: "gui_aramos@outlook.com",
		}

		user, err := repo.FindOne(context.TODO(), u.ID)

		assert.ErrorIs(t, primitive.ErrInvalidHex, err)
		assert.Empty(t, user)
	})

	mt.Run("should return error if find fails", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		user, err := repo.FindOne(context.TODO(), u.ID)

		assert.Empty(t, user)
		assert.Error(t, err)
	})

	mt.Run("should find one return empty user if not found", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))

		user, err := repo.FindOne(context.TODO(), u.ID)

		assert.NoError(t, err)
		assert.Empty(t, user)
	})
}

func TestUserRepository_FindByEmail(t *testing.T) {
	mtOpts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, mtOpts)
	defer mt.Close()

	mt.Run("should find one with success", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: u.ID},
			{Key: "email", Value: u.Email},
		}))

		user, err := repo.FindByEmail(context.TODO(), u.Email)

		assert.NoError(t, err)
		assert.Equal(t, u, user)
	})

	mt.Run("should return error if find fails", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		user, err := repo.FindByEmail(context.TODO(), u.Email)

		assert.Empty(t, user)
		assert.Error(t, err)
	})

	mt.Run("should find one return empty user if not found", func(mt *mtest.T) {
		repo := repository.NewUserRepository(mt.DB)

		u := &user.User{
			ID:    primitive.NewObjectID().Hex(),
			Email: "gui_aramos@outlook.com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))

		user, err := repo.FindByEmail(context.TODO(), u.Email)

		assert.NoError(t, err)
		assert.Empty(t, user)
	})
}
