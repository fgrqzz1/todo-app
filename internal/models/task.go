package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Done        bool `gorm:"default:false"`
	CreatedAt   time.Time
}
