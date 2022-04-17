package event

import (
	"fmt"
	_entities "project3/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	t.Run("TestGetEventsSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		events, err := eventUseCase.GetEvents()
		assert.Nil(t, err)
		assert.Equal(t, "Pengenalan Golang", events[0].Name)
	})
	t.Run("TestGetEventsError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		events, err := eventUseCase.GetEvents()
		assert.NotNil(t, err)
		assert.Nil(t, events)
	})
}

func TestGetEventByUserId(t *testing.T) {
	t.Run("TestGetEventsSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		events, err := eventUseCase.GetEventByUserId(1)
		assert.Nil(t, err)
		assert.Equal(t, "Pengenalan Golang", events[0].Name)
	})
	t.Run("TestGetEventsError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		events, err := eventUseCase.GetEventByUserId(1)
		assert.NotNil(t, err)
		assert.Nil(t, events)
	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("TestUpdateSuccess1", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Location: "Bandung"}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "Bandung", event.Location)
		assert.Equal(t, 100, event.Quota)
	})
	t.Run("TestUpdateSuccess2", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Quota: 100}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "Bandung", event.Location)
		assert.Equal(t, 100, event.Quota)
	})
	t.Run("TestUpdateError1", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Location: "Bandung"}, 1, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, rows)
		assert.NotEqual(t, "Bandung", event.Location)
		assert.NotEqual(t, 100, event.Quota)
	})
	t.Run("TestUpdateError2", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Quota: 100}, 1, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, rows)
		assert.NotEqual(t, "Bandung", event.Location)
		assert.NotEqual(t, 100, event.Quota)
	})
}

func TestCreateEvent(t *testing.T) {
	t.Run("TestCreateEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		err := eventUseCase.CreateEvent(1, _entities.Event{
			CategoryID:   1,
			Name:         "Pengenalan Golang",
			Host:         "Husnul Nawafil",
			Location:     "Jakarta",
			Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
			Quota:        200,
			Participants: uint(0),
		}, "unittest.png")
		assert.Nil(t, err)
	})
	t.Run("TestCreateEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		err := eventUseCase.CreateEvent(1, _entities.Event{
			CategoryID:   1,
			Name:         "Pengenalan Golang",
			Host:         "Husnul Nawafil",
			Location:     "Jakarta",
			Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
			Quota:        200,
			Participants: uint(0),
		}, "unittest.png")
		assert.NotNil(t, err)
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("TestDeleteEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		rows, err := eventUseCase.DeleteEvent(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})
	t.Run("TestDeleteEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		rows, err := eventUseCase.DeleteEvent(1, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, rows)
	})
}

func TestGetEventById(t *testing.T) {
	t.Run("TestGetEventByIdSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		event, err := eventUseCase.GetEventById(1)
		assert.Nil(t, err)
		assert.Equal(t, "Pengenalan Golang", event.Name)
		assert.Equal(t, "Jakarta", event.Location)
		assert.Equal(t, 200, event.Quota)
	})
	t.Run("TestGetEventByIdError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{})
		event, err := eventUseCase.GetEventById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, "Pengenalan Golang", event.Name)
		assert.NotEqual(t, "Jakarta", event.Location)
		assert.NotEqual(t, 200, event.Quota)
	})
}

// === mock success ===
type mockEventRepository struct{}

func (m mockEventRepository) GetEvents() ([]_entities.Event, error) {
	return []_entities.Event{
		{UserID: uint(1),
			CategoryID:   uint(1),
			Name:         "Pengenalan Golang",
			Host:         "Husnul Nawafil",
			Location:     "Jakarta",
			Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
			Quota:        200,
			Participants: uint(0),
		},
	}, nil
}

func (m mockEventRepository) GetEventByUserId(idToken int) ([]_entities.Event, error) {
	return []_entities.Event{
		{UserID: uint(1),
			CategoryID:   uint(1),
			Name:         "Pengenalan Golang",
			Host:         "Husnul Nawafil",
			Location:     "Jakarta",
			Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
			Quota:        200,
			Participants: uint(0),
		},
	}, nil
}

func (m mockEventRepository) UpdateEvent(event _entities.Event, event_ID, idToken int) (_entities.Event, int, error) {
	return _entities.Event{
		UserID:       uint(1),
		CategoryID:   uint(1),
		Name:         "Pengenalan Golang",
		Host:         "Husnul Nawafil",
		Location:     "Bandung",
		Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
		Quota:        100,
		Participants: uint(0),
	}, 1, nil
}

func (m mockEventRepository) CreateEvent(user_ID int, events _entities.Event, imageurl string) error {
	return nil
}

func (m mockEventRepository) DeleteEvent(event_ID, user_ID int) (int, error) {
	return 1, nil
}

func (m mockEventRepository) GetEventById(event_ID int) (_entities.Event, error) {
	return _entities.Event{
		UserID:       uint(1),
		CategoryID:   uint(1),
		Name:         "Pengenalan Golang",
		Host:         "Husnul Nawafil",
		Location:     "Jakarta",
		Details:      "Berkenalan dengan bahasa pemrograman yang cepat dan efisien",
		Quota:        200,
		Participants: uint(0),
	}, nil
}

// === mock error ===
type mockEventRepositoryError struct{}

func (m mockEventRepositoryError) GetEvents() ([]_entities.Event, error) {
	return nil, fmt.Errorf("error to get all events")
}
func (m mockEventRepositoryError) UpdateEvent(event _entities.Event, event_ID, idToken int) (_entities.Event, int, error) {
	return _entities.Event{}, 0, fmt.Errorf("error to update event")
}
func (m mockEventRepositoryError) CreateEvent(user_ID int, events _entities.Event, imageurl string) error {
	return fmt.Errorf("error to create event")
}
func (m mockEventRepositoryError) DeleteEvent(event_ID, user_ID int) (int, error) {
	return 0, fmt.Errorf("error to delete event")
}
func (m mockEventRepositoryError) GetEventById(event_ID int) (_entities.Event, error) {
	return _entities.Event{}, fmt.Errorf("error to get event by id")
}
func (m mockEventRepositoryError) GetEventByUserId(idToken int) ([]_entities.Event, error) {
	return nil, fmt.Errorf("error to get user's events")
}
