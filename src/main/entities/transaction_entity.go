package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	Checkin   time.Time      `gorm:"column:checkin;unique" json:"checkin"`
	Checkout  time.Time      `gorm:"column:checkout;nullable" json:"checkout"`
	UserID    uint           `gorm:"column:user_id" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"-"`
}

func (Transaction) TableName() string {
	return "transactions"
}
