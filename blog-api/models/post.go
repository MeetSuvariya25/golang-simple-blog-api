package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(100)" json:"title"`
	Description string `gorm"type:text" json:"description"`
	UserID      uint64 `gorm:"not null"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"User"`
}
