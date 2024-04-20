package api

import (
	"encoding/json"
	"fmt"
	"rideshare/foundation/pubsub"

	"github.com/google/uuid"
)

type service struct {
	pubsub *pubsub.PubSub
}

type DriverMatchingServiceConfig struct {
	BrokerURL string
	GroupID   string
}

func NewService(cfg DriverMatchingServiceConfig) (*service, error) {
	pubSub, err := pubsub.NewPubSub(pubsub.PubSubConfig{
		Brokers: cfg.BrokerURL,
		GroupID: cfg.GroupID,
	})
	if err != nil {
		return nil, err
	}

	return &service{pubsub: pubSub}, nil
}

func (s *service) StartConsuming() {
	errChan := make(chan error)

	go func() {
		for err := range errChan {
			fmt.Println("Error:", err)
		}
	}()

	events := s.pubsub.Consume("ride-requests", errChan)

	go func() {
		for e := range events {

			// Decode the event data field and store it in the ride request event type
			var rideRequest RideRequestEvent
			if err := json.Unmarshal(e.Data, &rideRequest); err != nil {
				errChan <- err
				continue
			}

			// Match the driver (replace with your matching logic)
			driverID := matchDriver(rideRequest.PickupLocation)

			// Prepare the driver assigned event
			driverAssignedEvent := DriverAssignedEvent{
				RideID:   rideRequest.RideID,
				DriverID: driverID,
			}
			// Convert event to JSON
			eventData, err := json.Marshal(driverAssignedEvent)
			if err != nil {
				errChan <- err
				continue
			}

			eventKey := uuid.New().String()

			// Publish to the topic driver-assignments
			if err := s.pubsub.Publish("driver-assignments", []byte(eventKey), eventData); err != nil {
				errChan <- err
			}
		}
	}()

}

func matchDriver(location Location) string {
	return "driver123"
}

func publishDriverAssignedEvent(event DriverAssignedEvent) {
	fmt.Printf("Published driver assigned event: %+v\n", event)
}
