package api

import (
	"encoding/json"
	"net/http"
	"rideshare/services/drivermatching/internal/service/api"
)

type service interface {
	StartConsuming()
}

type serivceHandler struct {
	service service
}

type DriverMatchingHandlerConfig struct {
	BrokerURL string
}

func NewHandler(cfg DriverMatchingHandlerConfig) (*serivceHandler, error) {

	service, err := api.NewService(api.DriverMatchingServiceConfig{
		BrokerURL: cfg.BrokerURL,
		GroupID:   "rideshare",
	})
	if err != nil {
		return nil, err
	}
	return &serivceHandler{service: service}, nil
}

func (h *serivceHandler) StartService() {

	h.service.StartConsuming()

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
