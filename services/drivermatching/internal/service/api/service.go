package api

import (
	"fmt"
	"rideshare/foundation/pubsub"
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

	messages := s.pubsub.Consume("ride-requests", errChan)

	go func() {
		for msg := range messages {
			fmt.Println("Message:", msg)
		}
	}()

}
