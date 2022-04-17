package event

import (
	_entities "project3/entities"
)

type EventRepositoryInterface interface {
	CreateEvent(user_ID int, events _entities.Event, imageurl string) error
	GetEvents() ([]_entities.Event, error)
	GetEventById(event_ID int) (_entities.Event, error)
	DeleteEvent(eventID, user_ID int) (int, error)
	UpdateEvent(event _entities.Event, event_ID, idToken int) (_entities.Event, int, error)
	GetEventByUserId(idToken int) ([]_entities.Event, error)
}
