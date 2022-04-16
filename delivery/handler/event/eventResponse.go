package event

import "time"

type GetEventResponse struct {
	ID       uint      `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Host     string    `json:"host" form:"host"`
	Date     time.Time `json:"date" form:"date"`
	Location string    `json:"location" form:"location"`
}
