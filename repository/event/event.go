package event

import (
	"errors"
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

func (er *EventRepository) CreateEvent(user_ID int, events _entities.Event, imageurl string) error {
	tx := er.DB.Exec("INSERT INTO events (user_id, category_id, name, host, date,location,details,quota, participants, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user_ID, events.CategoryID, events.Name, events.Host, events.Date, events.Location, events.Details, events.Quota, events.Participants, imageurl)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("unable to save event")
	}
	return nil

}

func (er *EventRepository) GetEvents() ([]_entities.Event, error) {
	var events []_entities.Event
	tx := er.DB.Preload("Attendees").Preload("Comment").Find(&events)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return events, nil
}
