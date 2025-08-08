package repository

import (
    "context"
    "errors"

    "github.com/google/uuid"
    "gorm.io/gorm"
    "vietceramics-apis/domain/model"
    repo "vietceramics-apis/domain/repository"
)

// postgresUserRepository implements the UserRepository interface using GORM and PostgreSQL.
type postgresUserRepository struct {
    db *gorm.DB
}

// NewUserRepository creates a new UserRepository backed by PostgreSQL.
func NewUserRepository(db *gorm.DB) repo.UserRepository {
    return &postgresUserRepository{db: db}
}

// GetByUsername retrieves a user by their username.
func (r *postgresUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
    var user model.User
    if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

// GetByID retrieves a user by their ID.
func (r *postgresUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
    var user model.User
    if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

// Create inserts a new user into the database.
func (r *postgresUserRepository) Create(ctx context.Context, user *model.User) error {
    if user.ID == uuid.Nil {
        user.ID = uuid.New()
    }
    return r.db.WithContext(ctx).Create(user).Error
}

// List returns all users.
func (r *postgresUserRepository) List(ctx context.Context) ([]*model.User, error) {
    var users []*model.User
    if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
