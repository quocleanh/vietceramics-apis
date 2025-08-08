package repository

import (
    "context"

    "github.com/google/uuid"
    "vietceramics-apis/domain/model"
)

// UserRepository defines the data access methods for the User entity.
// It is an abstraction over the underlying persistence mechanism.
type UserRepository interface {
    // GetByUsername returns a user by username or an error if not found.
    GetByUsername(ctx context.Context, username string) (*model.User, error)
    // GetByID returns a user by ID or an error if not found.
    GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
    // Create inserts a new user into the database.
    Create(ctx context.Context, user *model.User) error
    // List returns all users.
    List(ctx context.Context) ([]*model.User, error)
}
