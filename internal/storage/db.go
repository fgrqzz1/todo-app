package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo-app/internal/models"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connect with db: %w", err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return fmt.Errorf("Database schema migration error: %w", err)
	}
	return nil
}
