package usecase

import (
	"context"

	"github.com/guiaramos/umarket/cmd/user/domain/user"
	"github.com/guiaramos/umarket/pkg/apperror"
	"github.com/guiaramos/umarket/pkg/logger"
)

// CreateUserUseCase implements the logic for creatring an User.
type CreateUserUseCase struct {
	repo   user.Repository
	logger logger.Logger
}

// NewCreateUserUseCase creates a new Create User use case.
func NewCreateUserUseCase(repo user.Repository) CreateUserUseCase {
	logger := logger.NewLogger("create_user_usecase")

	return CreateUserUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u CreateUserUseCase) Execute(ctx context.Context, dto *CreateUserDTO) *apperror.AppError {
	u.logger.Info("CreateUserUserCase...")

	// Check if email is already in use.
	existingUser, err := u.repo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return ErrEmailAlreadyInUse
	}

	// Create the new user.
	id := u.repo.NewId()
	user := user.New(id, dto.Email)

	// Register the user.
	if err := user.Register(dto.Password); err != nil {
		return err
	}

	// Persist data.
	if err := u.repo.InsertOne(ctx, &user); err != nil {
		return err
	}

	return nil
}
