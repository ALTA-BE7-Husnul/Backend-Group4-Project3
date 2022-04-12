package event

import (
	_entities "project3/entities"

	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		DB: db,
	}
}

func (er *EventRepository) GetEvents() ([]_entities.Event, error) {
	var events []_entities.Event

	return events, nil
}
