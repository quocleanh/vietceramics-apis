package model

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User represents an account in the system. It is stored in the users
// table and used for authentication and authorization. Passwords should
// be stored as bcrypt hashes.
type User struct {
    ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
    Username  string         `gorm:"uniqueIndex;not null" json:"username"`
    Password  string         `json:"-"` // hashed password stored securely
    FullName  string         `json:"full_name"`
    Email     string         `json:"email"`
    Role      string         `json:"role"` // SA, USER, etc.
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
