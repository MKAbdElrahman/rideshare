package api

import (
	"encoding/json"
	"rideshare/foundation/pubsub"
	"time"

	"github.com/google/uuid"
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

	// save the ride request in the database and return the ride id
	rideID := uuid.New().String()

	eventData, err := json.Marshal(RideRequestEvent{
		RideID:          rideID,
		UserID:          request.UserID,
		PickupLocation:  request.PickupLocation,
		DropoffLocation: request.DropoffLocation,
	})
	if err != nil {
		return Ride{}, err
	}
	
	err = s.pubsub.Publish("ride-requests", []byte(rideID), eventData)
	if err != nil {
		return Ride{}, err
	}
	return Ride{
		RideID:          rideID,
		UserID:          request.UserID,
		PickupLocation:  request.PickupLocation,
		DropoffLocation: request.DropoffLocation,
		Status:          "RideRequestCreated",
		CreatedAt:       time.Now(),
	}, nil
}
