package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohexists/event-broker-svc/config"
	"github.com/bohexists/event-broker-svc/domain"
	"github.com/bohexists/event-broker-svc/internal/adapters/nats"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create NATS publisher
	natsPublisher, err := nats.NewNatsPublisher(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Failed to create NATS publisher: %v", err)
	}

	// Create an example event to publish
	event := domain.NewEvent("task.status.update", []byte(`{"task_id": 123, "status": "completed"}`))

	// Publish the event
	err = natsPublisher.Publish(event.Subject, event.Data)
	if err != nil {
		log.Fatalf("Failed to publish event: %v", err)
	}

	log.Println("Event published successfully!")

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down...")
}
