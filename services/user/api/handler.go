package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rideshare/services/user/internal/userservice/api"
	"strconv"
)

type userService interface {
	Register(param api.UserRegistrationParam) error
	GetUserByID(id int) (api.PublicUser, error)
}

type userHandler struct {
	userService userService
}

type HandlerConfig struct {
	KafkaBootstrapServer string
	Topic                string
}

func NewUserHandler(cfg HandlerConfig) (*userHandler, error) {

	us, err := api.NewUserService(api.ServiceConfig{})
	if err != nil {
		return nil, err
	}
	return &userHandler{userService: us}, nil
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registrationParams api.UserRegistrationParam

	err := json.NewDecoder(r.Body).Decode(&registrationParams)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userService.Register(registrationParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User registration successful!")
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing user ID parameter", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}
}
