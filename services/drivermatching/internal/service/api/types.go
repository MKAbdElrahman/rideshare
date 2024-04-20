package api

import "time"

type Thing struct{}

type DriverAssignedEvent struct {
	RideID   string `json:"ride_id"`
	DriverID string `json:"driver_id"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RideRequestEvent struct {
	RideID          string    `json:"ride_id"`
	UserID          string    `json:"user_id"`
	PickupLocation  Location  `json:"pickup_location"`
	DropoffLocation Location  `json:"dropoff_location"`
	CreatedAt       time.Time `json:"created_at"`
}
