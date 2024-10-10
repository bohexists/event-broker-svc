package domain

// Event represents the domain model for an event that can be published.
type Event struct {
	ID      string
	Subject string
	Data    []byte
}

// NewEvent creates a new event.
func NewEvent(subject string, data []byte) *Event {
	return &Event{
		Subject: subject,
		Data:    data,
	}
}
