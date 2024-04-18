package api

import (
	"rideshare/services/service/internal/service/internal/db"
)

type thingRepository interface {
	CreateThing(us db.Thing) error
	GetThingByID(thingID int) (db.Thing, error)
}

type publisherI interface {
	Publish(string) error
}

type service struct {
	repo      thingRepository
	publisher publisherI
}

type ServiceConfig struct {
}

func NewService(cfg ServiceConfig) (*service, error) {

	return &service{}, nil
}

func (s *service) Action() error {

	return nil
}
