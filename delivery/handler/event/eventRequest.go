package event

import (
	"gorm.io/gorm"
)

type EventRequest struct {
	gorm.Model
	UserID       uint   `json:"user_id" form:"user_id"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	Name         string `json:"name" form:"name"`
	Host         string `json:"host" form:"host"`
	Date         string `json:"date" form:"date"`
	Location     string `json:"location" form:"location"`
	Details      string `json:"details" form:"details"`
	Quota        int    `json:"quota" form:"quota"`
	Participants uint   `json:"participants" form:"participants"`
	Image        string `json:"image" form:"image"`
}
