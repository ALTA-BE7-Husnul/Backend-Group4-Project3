package event

import (
	_entities "project3/entities"
)

type EventRepositoryInterface interface {
	CreateEvent(user_ID int, events _entities.Event, imageurl string) error
	GetEvents() ([]_entities.Event, error)
	DeleteEvent(eventID, user_ID int) (int, error)
}
