package attendees

import (
	_entities "project3/entities"
)

type AttendeesUseCaseInterface interface {
	CreateAttendees(request _entities.Attendees) (_entities.Attendees, error)
}