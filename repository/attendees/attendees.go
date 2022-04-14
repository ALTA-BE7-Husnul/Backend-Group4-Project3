package attendees

import (
	_entities "project3/entities"

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

func (ur *AttendeesRepository) CreateAttendees(request _entities.Attendees) (_entities.Attendees, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}
	return request, nil
}
