package api

import (
	"encoding/json"
	"net/http"
	"rideshare/services/riderequest/internal/service/api"
)

type service interface {
	CreateRide(requst api.RideRequestParams) (api.Ride, error)
}

type serivceHandler struct {
	service service
}

type RideRequestHandlerConfig struct {
	BrokerURL string
}

func NewHandler(cfg RideRequestHandlerConfig) (*serivceHandler, error) {

	service, err := api.NewService(api.RideRequestServiceConfig{
		BrokerURL: cfg.BrokerURL,
	})
	if err != nil {
		return nil, err
	}
	return &serivceHandler{service: service}, nil
}

func (h *serivceHandler) HandleCreateRide(w http.ResponseWriter, r *http.Request) {
	// 1. Parse request body
	var request api.RideRequestParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Validate request data
	// err = h.service.ValidateRideRequest(request)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// 3. Call service to create ride request
	createdRide, err := h.service.CreateRide(request)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 4. Respond with created ride details (limited to ID and status for this scope)
	response := struct {
		Ride api.Ride `json:"ride"`
	}{

		Ride: createdRide,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (h *serivceHandler) HandleGetRide(w http.ResponseWriter, r *http.Request) {
}

func (h *serivceHandler) HandleUpdateRide(w http.ResponseWriter, r *http.Request) {
}

func (h *serivceHandler) HandleDeleteRide(w http.ResponseWriter, r *http.Request) {
}

func (h *serivceHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {

	status := map[string]string{"status": "ok"}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
