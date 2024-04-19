package main

import (
	"net/http"
	dirverMatchingServiceAPI "rideshare/services/drivermatching/api"
	rideRequestServiceAPI "rideshare/services/riderequest/api"

	userServiceAPI "rideshare/services/user/api"
)

func main() {

	mux := http.NewServeMux()

	userHandler, err := userServiceAPI.NewUserHandler(userServiceAPI.HandlerConfig{
		KafkaBootstrapServer: "localhost:9092",
		Topic:                "users",
	})

	if err != nil {
		panic(err)
	}

	mux.HandleFunc("POST /api/v1/user/register", userHandler.Register)
	mux.HandleFunc("GET /api/v1/user/{id}", userHandler.GetUserByID)
	mux.HandleFunc("GET /api/v1/user/health", userHandler.HandleHealthCheck)

	rideRequestHandler, err := rideRequestServiceAPI.NewHandler(rideRequestServiceAPI.RideRequestHandlerConfig{
		BrokerURL: "localhost:9092",
	})

	if err != nil {
		panic(err)
	}

	mux.HandleFunc("POST /api/v1/ride/create", rideRequestHandler.HandleCreateRide)
	mux.HandleFunc("GET /api/v1/ride/health", rideRequestHandler.HandleHealthCheck)

	driverMatchingHandler, err := dirverMatchingServiceAPI.NewHandler(dirverMatchingServiceAPI.DriverMatchingHandlerConfig{
		BrokerURL: "localhost:9092",
	})
	if err != nil {
		panic(err)
	}

	driverMatchingHandler.StartService()

	
	http.ListenAndServe(":3000", mux)
}
