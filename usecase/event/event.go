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
func (euc *EventUseCase) GetEventById(event_ID int) (_entities.Event, error) {
	event, err := euc.eventRepository.GetEventById(event_ID)
	return event, err
}

func (euc *EventUseCase) DeleteEvent(event_ID, user_ID int) (int, error) {
	rows, err := euc.eventRepository.DeleteEvent(event_ID, user_ID)
	return rows, err
}

func (euc *EventUseCase) UpdateEvent(event _entities.Event, event_ID, idToken int) (_entities.Event, int, error) {
	data, rows, err := euc.eventRepository.UpdateEvent(event, event_ID, idToken)
	return data, rows, err
}
