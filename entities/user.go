package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Comment  []Comment `gorm:"foreignKey:UserID;references:ID"`
}
