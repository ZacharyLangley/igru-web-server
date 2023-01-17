package proxy

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func RegisterProxy(r *mux.Router, cfg config.GRPC, serviceName string) error {
	dstURL, err := url.Parse(cfg.Address)
	if err != nil {
		return err
	}
	p := &GRPC{
		URL: dstURL,
	}
	routeName := "/" + serviceName + "/"
	zap.L().Debug("registered proxy", zap.String("path", routeName), zap.String("url", dstURL.String()))
	r.PathPrefix(routeName).Handler(p)
	return nil
}

type GRPC struct {
	URL *url.URL
}

func (h GRPC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Debug("proxying http", zap.String("host", h.URL.Host), zap.String("scheme", h.URL.Scheme))
	if r.Method == http.MethodConnect {
		h.handleTunneling(w, r)
	} else {
		h.handleHTTP(w, r)
	}
}

func (h GRPC) handleTunneling(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", h.URL.Host, 10*time.Second)
	if err != nil {
		zap.L().Error("failed to dial upstream", zap.String("upstream", h.URL.Host))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		zap.L().Warn("hijacking not supported", zap.String("upstream", h.URL.Host))
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		zap.L().Error("hijacker failure", zap.String("upstream", h.URL.Host), zap.Error(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func (h GRPC) handleHTTP(w http.ResponseWriter, req *http.Request) {
	req.URL.Scheme = h.URL.Scheme
	req.URL.Host = h.URL.Host
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		zap.L().Error("failed to dial upstream", zap.String("upstreamHost", h.URL.Host), zap.String("upstreamScheme", h.URL.Scheme), zap.String("requestScheme", req.URL.Scheme))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
