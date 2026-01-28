package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo-app/internal/models"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Connect db: %w", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, fmt.Errorf("Auto migrate: %w", err)
	}

	return db, nil
}
