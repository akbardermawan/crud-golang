package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id"  gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deleted_at" `
}