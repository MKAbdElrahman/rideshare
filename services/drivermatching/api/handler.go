package api

import (
	"net/http"
	"rideshare/services/service/internal/service/api"
)

// this interface to constrain what the user handler needs and also works as a contract
type service interface {
}

type serivceHandler struct {
	service service
}

type HandlerConfig struct {
	KafkaBootstrapServer string
	Topic                string
}


// The service handler creates the concete implmentation as i done't want  the caller to have the responsibility of  injecting the dependancy
func NewServiceHandler(cfg HandlerConfig) (*serivceHandler, error) {

	s, err := api.NewService(api.ServiceConfig{})
	if err != nil {
		return nil, err
	}
	return &serivceHandler{service: s}, nil
}

func (h *serivceHandler) HandleSomething(w http.ResponseWriter, r *http.Request) {
}
