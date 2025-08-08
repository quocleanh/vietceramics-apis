package repository

import (
    "context"

    "github.com/google/uuid"
    "vietceramics-apis/domain/model"
)

// PermissionRepository defines the data access methods for Permission entities.
type PermissionRepository interface {
    // GetPermissionsForUser returns all permissions for a given user.
    GetPermissionsForUser(ctx context.Context, userID uuid.UUID) ([]*model.Permission, error)
    // IsAllowed checks if a user is allowed to access the given endpoint with the given method.
    IsAllowed(ctx context.Context, userID uuid.UUID, endpoint, method string) (bool, error)
}
