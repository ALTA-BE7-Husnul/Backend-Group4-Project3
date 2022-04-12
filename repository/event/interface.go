package event

import (
	_entities "project3/entities"
)

type EventRepositoryInterface interface {
	GetEvents() ([]_entities.Event, error)
}
