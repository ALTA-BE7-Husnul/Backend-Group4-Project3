package attendees

import (
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
