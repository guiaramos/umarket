package repository

import (
	"context"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/pkg/apperror"
	"github.com/guiaramos/umarket/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll   *mongo.Collection
	logger logger.Logger
}

// NewUserRepository creates new user mongodb repository
func NewUserRepository(mongo *mongo.Database) user.Repository {
	coll := mongo.Collection("users")
	logger := logger.NewLogger("user_repository")

	return &userRepository{coll: coll, logger: logger}
}

// NewId generates a new mongodb id
func (r userRepository) NewId() string {
	return primitive.NewObjectID().Hex()
}

// InsertOne persist new user to mongodb
func (r userRepository) InsertOne(ctx context.Context, u *user.User) *apperror.AppError {
	result, err := r.coll.InsertOne(ctx, u)
	if err != nil {
		return apperror.NewInternalServerError(err.Error())
	}

	r.logger.InfoWithFields("New user created with the following id: ", logger.LoggerFields{"id:": result.InsertedID})

	return nil
}

// UpdateOne saves current user changes to mongodb
func (r userRepository) UpdateOne(ctx context.Context, u *user.User) *apperror.AppError {
	id, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return apperror.NewBadRequest(err.Error())
	}

	update := bson.D{{Key: "$set", Value: u}}

	result, err := r.coll.UpdateByID(ctx, id, update)
	if err != nil {
		return apperror.NewInternalServerError(err.Error())
	}

	r.logger.InfoWithFields("document(s) was/were updated.", logger.LoggerFields{
		"MatchedCount":  result.MatchedCount,
		"ModifiedCount": result.ModifiedCount,
		"UpsertedCount": result.UpsertedCount,
		"UpsertedID":    result.UpsertedID,
	})

	return nil
}

// FindOne searchs for one user from mongodb
func (r userRepository) FindOne(ctx context.Context, id string) (*user.User, *apperror.AppError) {
	u := &user.User{}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, apperror.NewBadRequest(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	if err = r.coll.FindOne(ctx, filter).Decode(u); err != nil {
		return u, apperror.NewInternalServerError(err.Error())
	}

	if u.ID != "" {
		r.logger.InfoWithFields("Found a user in the collection with the id: ", logger.LoggerFields{"id": u.ID})
	} else {
		r.logger.InfoWithFields("No users found with the id: ", logger.LoggerFields{"id": u.ID})
	}

	return u, nil
}

// FindByEmail searchs for user with passed email
func (r userRepository) FindByEmail(ctx context.Context, email user.EmailAddress) (*user.User, *apperror.AppError) {
	u := &user.User{}

	filter := bson.D{{Key: "email", Value: email}}

	err := r.coll.FindOne(ctx, filter).Decode(u)
	if err != nil {
		return u, apperror.NewInternalServerError(err.Error())
	}

	if u.Email != "" {
		r.logger.InfoWithFields("Found a user in the collection with the email: ", logger.LoggerFields{"email": u.Email})
	} else {
		r.logger.InfoWithFields("No users found with the email: ", logger.LoggerFields{"email": u.Email})
	}

	return u, nil
}
