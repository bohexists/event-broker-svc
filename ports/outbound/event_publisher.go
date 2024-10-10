package outbound

// EventPublisher defines the contract for publishing events.
type EventPublisher interface {
	Publish(subject string, data []byte) error
}
