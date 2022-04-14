package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID   uint      `json:"user_id" form:"user_id"`
	CategoryID uint  `json:"category_id" form:"category_id"`
	Name     string    `json:"name" form:"name"`
	Host     string    `json:"host" form:"host"`
	Date     time.Time `json:"date" form:"date"`
	Location string    `json:"location" form:"location"`
	Details  string    `json:"details" form:"details"`
	Quota    int       `json:"quota" form:"quota"`
	Participants uint  `gorm:"not null" json:"participants" form:"participants"`
	Image    string    `json:"image" form:"image"`
	Comment []Comment  ` gorm:"foreignKey:EventID;references:ID" json:"comment" form:"comment"`
	Attendees []Attendees  `gorm:"foreignKey:EventID;references:ID" json:"attendees" form:"attendees"`
}
