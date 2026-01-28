package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo-app/internal/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Не удалось подключиться к БД: %v", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		fmt.Printf("Не удалось выполнить миграцию: %v", err)
	}

	return db
}
