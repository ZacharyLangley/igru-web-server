package proxy

import (
	"net/url"
)

type GRPC struct {
	URL *url.URL
}

func (g GRPC) GetURL() *url.URL {
	return g.URL
}
