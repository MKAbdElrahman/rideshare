package api

import "time"

type RideRequestParams struct {
	UserID         string   `json:"user_id"`
	PickupLocation Location `json:"pickup_location"` // Pickup location details
	Destination    Location `json:"destination"`     // Destination location details
}

type Ride struct {
	ID             string    `json:"id"`              // Unique identifier for the ride request
	UserID         string    `json:"user_id"`         // ID of the user who requested the ride
	PickupLocation Location  `json:"pickup_location"` // Pickup location details
	Destination    Location  `json:"destination"`     // Destination location details
	Status         string    `json:"status"`          // Current status of the ride request (e.g., "Created", "Searching for driver", "Driver assigned")
	CreatedAt      time.Time `json:"created_at"`      // Timestamp of ride request creation
}

type Location struct {
	Latitude  float64 `json:"latitude"`  // Geographic latitude
	Longitude float64 `json:"longitude"` // Geographic longitude
}

type RideRequestEvent struct {
	ID      string            `json:"id"`
	Request RideRequestParams `json:"request"`
}
