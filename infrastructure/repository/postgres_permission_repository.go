package repository

import (
    "context"
    "errors"

    "github.com/google/uuid"
    "gorm.io/gorm"
    "vietceramics-apis/domain/model"
    repo "vietceramics-apis/domain/repository"
)

// postgresPermissionRepository implements the PermissionRepository interface using GORM.
type postgresPermissionRepository struct {
    db *gorm.DB
}

// NewPermissionRepository creates a new PermissionRepository backed by PostgreSQL.
func NewPermissionRepository(db *gorm.DB) repo.PermissionRepository {
    return &postgresPermissionRepository{db: db}
}

// GetPermissionsForUser returns all permissions for the given user.
func (r *postgresPermissionRepository) GetPermissionsForUser(ctx context.Context, userID uuid.UUID) ([]*model.Permission, error) {
    var perms []*model.Permission
    if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&perms).Error; err != nil {
        return nil, err
    }
    return perms, nil
}

// IsAllowed checks if a user has an allowed permission for the given endpoint and method.
func (r *postgresPermissionRepository) IsAllowed(ctx context.Context, userID uuid.UUID, endpoint, method string) (bool, error) {
    var perm model.Permission
    if err := r.db.WithContext(ctx).Where("user_id = ? AND endpoint = ? AND method = ?", userID, endpoint, method).First(&perm).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
        return false, err
    }
    return perm.Allowed, nil
}
