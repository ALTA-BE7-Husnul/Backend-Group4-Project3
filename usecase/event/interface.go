package event

import (
	_entities "project3/entities"
)

type EventUseCaseInterface interface {
	CreateEvent(user_ID int, events _entities.Event, imageurl string) error
	GetEvents() ([]_entities.Event, error)
	DeleteEvent(event_ID, user_ID int) (int, error)
	UpdateEvent(event _entities.Event, event_ID, idToken int, imageurl string) (_entities.Event, int, error)
}
