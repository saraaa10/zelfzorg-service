package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey;column:id" json:"id"`
	Name         string         `gorm:"column:name" json:"name"`
	Email        string         `gorm:"column:email;unique" json:"email"`
	Username     string         `gorm:"column:username" json:"username"`
	Password     string         `gorm:"column:password" json:"password"`
	UsetTypeId   uint           `gorm:"column:user_type_id" json:"-"`
	UserType     UserType       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_type"`
	Transactions []Transaction  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transactions"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (User) TableName() string {
	return "users"
}
