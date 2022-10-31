package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserType struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (UserType) TableName() string {
	return "user_types"
}
