package attendees

import (
	_entities "project3/entities"
)

type AttendeesRepositoryInterface interface {
	CreateAttendees(request _entities.Attendees) (_entities.Attendees, error)
	
}