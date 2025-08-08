package model

import (
    "github.com/google/uuid"
)

// Permission links a user to a specific API endpoint and HTTP method with an allow/deny rule.
// This can be used in addition to role-based checks for more fine‑grained access control.
type Permission struct {
    ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    Endpoint string    `gorm:"not null" json:"endpoint"`
    Method   string    `gorm:"not null" json:"method"`
    Allowed  bool      `gorm:"default:false" json:"allowed"`
}
