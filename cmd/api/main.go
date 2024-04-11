package main

import (
	"net/http"
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

	http.ListenAndServe(":3000", mux)
}
