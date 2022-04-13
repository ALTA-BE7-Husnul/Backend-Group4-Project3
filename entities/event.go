package entities

import "time"

type Event struct {
	UserID   uint      `json:"user_id" form:"user_id"`
	Category string    `json:"category" form:"category"`
	Name     string    `json:"name" form:"name"`
	Host     string    `json:"host" form:"host"`
	Date     time.Time `json:"date" form:"date"`
	Location string    `json:"location" form:"location"`
	Details  string    `json:"details" form:"details"`
	Quota    int       `json:"quota" form:"quota"`
	Image    string    `json:"image" form:"image"`
}
