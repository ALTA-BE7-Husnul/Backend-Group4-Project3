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
	events.Image = imageurl
	events.UserID = uint(user_ID)
	// tx := er.DB.Exec("INSERT INTO events (user_id, category_id, name, host, date,location,details,quota, participants, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user_ID, events.CategoryID, events.Name, events.Host, events.Date, events.Location, events.Details, events.Quota, events.Participants, imageurl)
	tx := er.DB.Save(&events)
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

func (er *EventRepository) GetEventById(event_ID int) (_entities.Event, error) {
	var event _entities.Event
	tx := er.DB.Preload("Attendees").Preload("Comment").Where("id = ?", event_ID).Find(&event)
	if tx.Error != nil {
		return _entities.Event{}, tx.Error
	}
	return event, nil
}

func (er *EventRepository) DeleteEvent(event_ID, user_ID int) (int, error) {
	var events []_entities.Event
	tx := er.DB.Where("id = ?", event_ID).Where("user_id = ?", user_ID).Delete(&events)
	if tx.Error != nil {
		return 0, tx.Error
	}
	rows := tx.RowsAffected
	if rows == 0 {
		return 0, tx.Error
	}
	return int(rows), nil
}

func (er *EventRepository) UpdateEvent(event _entities.Event, event_ID, idToken int) (_entities.Event, int, error) {
	if event.Location != "" {
		tx := er.DB.Model(&_entities.Event{}).Where("id = ?", event_ID).Where("user_id = ?", idToken).Update("location", event.Location)
		if tx.Error != nil {
			return event, 1, tx.Error
		}
		if tx.RowsAffected == 0 {
			return event, 2, tx.Error
		}
	}
	if event.Quota > 0 {
		tx := er.DB.Model(&_entities.Event{}).Where("id = ?", event_ID).Where("user_id = ?", idToken).Update("quota", event.Quota)
		if tx.Error != nil {
			return event, 1, tx.Error
		}
		if tx.RowsAffected == 0 {
			return event, 2, tx.Error
		}
	}
	return event, 0, nil
}
