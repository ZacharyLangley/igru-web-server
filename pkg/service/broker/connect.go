package broker

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/proto/broker/v1/brokerv1connect"
)

func New() *Service {
	return &Service{}
}

type Service struct {
	brokerv1connect.UnimplementedBrokerServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.Handle(brokerv1connect.NewBrokerServiceHandler(s))
}
