package nats

import (
	"github.com/nats-io/nats.go"

	"log"
)

// NatsPublisher is an implementation of the EventPublisher interface for NATS.
type NatsPublisher struct {
	Connection *nats.Conn
}

// NewNatsPublisher creates a new instance of NatsPublisher with a NATS connection.
func NewNatsPublisher(natsURL string) (*NatsPublisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsPublisher{Connection: nc}, nil
}

// Publish publishes a message to a NATS subject.
func (n *NatsPublisher) Publish(subject string, data []byte) error {
	err := n.Connection.Publish(subject, data)
	if err != nil {
		log.Printf("Error publishing message to NATS: %v", err)
		return err
	}
	log.Printf("Message published to NATS successfully on subject: %s", subject)
	return nil
}
