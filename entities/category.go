package entities

type Category struct {
	ID           uint    `gorm:"primarykey"`
	CategoryName string  `gorm:"not null" json:"category_name" form:"category_name"`
	Event        []Event `gorm:"foreignKey:CategoryID;references:ID"`
}
