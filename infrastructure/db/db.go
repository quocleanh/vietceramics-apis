package db

import (
    "fmt"

    "vietceramics-apis/config"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// NewPostgresDB connects to a Postgres database using gorm and returns the *gorm.DB instance.
// It also performs automatic migrations for models defined in the domain layer. In this
// example we leave migrations to the caller to keep things simple.
func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
    // Use gorm's postgres driver to open a connection.
    db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("could not connect to database: %w", err)
    }
    // In a production system you would automatically migrate your schema here.
    return db, nil
}
