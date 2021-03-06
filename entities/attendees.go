package entities

import "gorm.io/gorm"

type Attendees struct {
	gorm.Model
	EventID uint  `gorm:"not null" json:"event_id" form:"event_id"`
	UserID  uint  `gorm:"not null" json:"user_id" form:"user_id"`
	User    User  `gorm:"foreignKey:UserID;references:ID"`
	Event   Event `gorm:"foreignKey:EventID;references:ID"`
}
