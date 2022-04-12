package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	EventID      uint   `gorm:"not null" json:"event_id" form:"event_id"`
	UserID       uint   `gorm:"not null" json:"user_id" form:"user_id"`
	Comment     string `gorm:"not null" json:"comment" form:"comment"`
}