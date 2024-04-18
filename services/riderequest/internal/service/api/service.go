package api

import (
	"encoding/json"
	"rideshare/foundation/pubsub"
)

type service struct {
	pubsub *pubsub.PubSub
}

type RideRequestServiceConfig struct {
	BrokerURL string
}

func NewService(cfg RideRequestServiceConfig) (*service, error) {
	pubSub, err := pubsub.NewPubSub(pubsub.PubSubConfig{
		Brokers: cfg.BrokerURL,
	})
	if err != nil {
		return nil, err
	}

	return &service{pubsub: pubSub}, nil
}

func (s *service) CreateRide(request RideRequestParams) (Ride, error) {
	eventID := generateUniqueID()

	event := RideRequestEvent{
		ID:      eventID,
		Request: request,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return Ride{}, err
	}
	err = s.pubsub.Publish("ride-requests", data)
	if err != nil {
		return Ride{}, err
	}
	return Ride{
		ID:             eventID,
		UserID:         request.UserID,
		PickupLocation: request.PickupLocation,
		Destination:    request.Destination,
		Status:         "RideRequestCreated",
	}, nil // Replace with actual ride creation logic and return value
}

func generateUniqueID() string {
	return "placeholder_id"
}