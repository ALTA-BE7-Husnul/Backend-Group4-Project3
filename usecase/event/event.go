package event

import (
	_entities "project3/entities"
	_eventRepository "project3/repository/event"
)

type EventUseCase struct {
	eventRepository _eventRepository.EventRepositoryInterface
}

func NewEventUseCase(eventRepo _eventRepository.EventRepositoryInterface) EventUseCaseInterface {
	return &EventUseCase{
		eventRepository: eventRepo,
	}
}

func (euc *EventUseCase) CreateEvent(user_ID int, events _entities.Event, imageurl string) error {
	err := euc.eventRepository.CreateEvent(user_ID, events, imageurl)
	return err
}

func (euc *EventUseCase) GetEvents() ([]_entities.Event, error) {
	events, err := euc.eventRepository.GetEvents()
	return events, err
}
