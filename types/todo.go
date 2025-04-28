package types

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	Task        string
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}
