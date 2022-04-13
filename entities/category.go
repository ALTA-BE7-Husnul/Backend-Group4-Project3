package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string `json:"category" form:"category"`
}
