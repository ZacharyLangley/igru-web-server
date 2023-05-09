package proxy

import (
	"net/url"
)

type HTTP struct {
	URL *url.URL
}

func (h HTTP) GetURL() *url.URL {
	return h.URL
}
