package event

import (
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
}

func TestUpdateEvent(t *testing.T) {
	t.Run("TestUpdateSuccess1", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Location: "Bandung"}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "Bandung", event.Location)
	})
	t.Run("TestUpdateSuccess2", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{Quota: 100}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "Bandung", event.Location)
		assert.Equal(t, 100, event.Quota)
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
}

func TestDeleteEvent(t *testing.T) {
	t.Run("TestDeleteEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{})
		rows, err := eventUseCase.DeleteEvent(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
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
