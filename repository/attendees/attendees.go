package attendees

import (
	"errors"
	_entities "project3/entities"

	"fmt"

	"gorm.io/gorm"
)

type AttendeesRepository struct {
	DB *gorm.DB
}

func NewAttendeesRepository(db *gorm.DB) *AttendeesRepository {
	return &AttendeesRepository{
		DB: db,
	}
}

func (ur *AttendeesRepository) CreateAttendees(request _entities.Attendees) (_entities.Attendees, int, error) {
	var event _entities.Event

	tx := ur.DB.Where("id = ?", request.EventID).Find(&event)

	if tx.RowsAffected == 0 {
		return _entities.Attendees{}, 6, fmt.Errorf("not found")
	}

	if tx.Error != nil {
		return _entities.Attendees{}, 1, fmt.Errorf("fail to read event")
	}

	if event.Participants == uint(event.Quota) {
		return _entities.Attendees{}, 2, fmt.Errorf("quota full")
	}

	var attendees []_entities.Attendees
	sx := ur.DB.Where("event_id = ?", request.EventID).Where("user_id = ?", request.UserID).Find(&attendees)

	if sx.RowsAffected > 0 {
		return _entities.Attendees{}, 3, fmt.Errorf("you have joined")
	}

	if sx.Error != nil {
		return _entities.Attendees{}, 4, fmt.Errorf("fail to read attendees")
	}

	yx := ur.DB.Save(&request)
	ur.DB.Exec("UPDATE events SET participants = ? WHERE id = ?", gorm.Expr("participants + ?", 1), request.EventID)
	if yx.Error != nil {
		return request, 5, yx.Error
	}
	return request, 0, nil
}

func (ur *AttendeesRepository) GetAttendees(request _entities.Attendees) ([]_entities.Attendees, error) {
	var attendees []_entities.Attendees
	tx := ur.DB.Preload("User").Where("event_id = ?", request.EventID).Find(&attendees)

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return attendees, nil
}

func (ur *AttendeesRepository) GetEventsByUserId(user_ID int) ([]_entities.Attendees, error) {
	var attendees []_entities.Attendees
	tx := ur.DB.Preload("Event").Where("user_id = ?", user_ID).Find(&attendees)
	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return attendees, nil
}

func (ur *AttendeesRepository) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	var attendees _entities.Attendees
	tx := ur.DB.Where("event_id = ?", idEvent).Where("user_id = ?", idToken).Unscoped().Delete(&attendees)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}

	ur.DB.Exec("UPDATE events SET participants = ? WHERE id = ?", gorm.Expr("participants - ?", 1), idEvent)

	return uint(tx.RowsAffected), nil
}
