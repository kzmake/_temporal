package workflow

// BookTripTaskQueue ...
const BookTripTaskQueue = "BOOK_TRIP_TASK_QUEUE"

// BaseEvent ...
type BaseEvent struct {
	ID string `json:"id"`
}

// BookedEvent ...
type BookedEvent struct {
	BaseEvent

	Name        string `json:"name"`
	Destination string `json:"destination"`
}

// NewBookedEvent ...
func NewBookedEvent(id, name, destination string) BookedEvent {
	return BookedEvent{
		BaseEvent:   BaseEvent{ID: id},
		Name:        name,
		Destination: destination,
	}
}
