package connect

import "net/http"

type Service interface {
	Register(*http.ServeMux)
}

func CreateMux(services ...Service) *http.ServeMux {
	mux := http.NewServeMux()
	for _, svc := range services {
		svc.Register(mux)
	}
	return mux
}
