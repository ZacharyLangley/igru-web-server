package proxy

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type HTTP struct {
	URL *url.URL
}

func (h HTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		h.handleTunneling(w, r)
	} else {
		h.handleHTTP(w, r)
	}
}

func (h HTTP) handleTunneling(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", h.URL.Host, 10*time.Second)
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
	go transfer(dest_conn, clientConn)
	go transfer(clientConn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	if _, err := io.Copy(destination, source); err != nil {
		zap.L().Error("transfer copy error", zap.Error(err))
	}
}
func (h HTTP) handleHTTP(w http.ResponseWriter, req *http.Request) {
	req.URL.Scheme = h.URL.Scheme
	req.URL.Host = h.URL.Host
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		zap.L().Error("failed to dial upstream", zap.String("upstream", h.URL.Host))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		zap.L().Error("copy error", zap.Error(err))
	}
}
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
