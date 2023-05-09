package proxy

import (
	"fmt"
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
		return fmt.Errorf("failed to parse address: %w", err)
	}
	p := NewTunnel("grpc", dstURL)
	routeName := "/" + serviceName + "/"
	zap.L().Debug("registered proxy", zap.String("path", routeName), zap.String("url", dstURL.String()))
	r.PathPrefix(routeName).Handler(p)
	return nil
}

func NewTunnel(name string, protocol *url.URL) *Tunnel {
	return &Tunnel{
		protocol:  name,
		targetURL: protocol,
	}
}

type Tunnel struct {
	protocol  string
	targetURL *url.URL
}

func (t Tunnel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Debug("proxying", zap.String("protocol", t.protocol), zap.String("host", t.targetURL.Host), zap.String("scheme", t.targetURL.Scheme))
	if r.Method == http.MethodConnect {
		t.handleTunneling(w, r)
	} else {
		t.handleHTTP(w, r)
	}
}

func (t Tunnel) handleTunneling(w http.ResponseWriter, _ *http.Request) {
	destConn, err := net.DialTimeout("tcp", t.targetURL.Host, 10*time.Second)
	if err != nil {
		zap.L().Error("failed to dial upstream", zap.String("upstream", t.targetURL.Host))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		zap.L().Warn("hijacking not supported", zap.String("upstream", t.targetURL.Host))
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		zap.L().Error("hijacker failure", zap.String("upstream", t.targetURL.Host), zap.Error(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	if _, err := io.Copy(destination, source); err != nil {
		zap.L().Error("transfer copy error", zap.Error(err))
	}
}
func (t Tunnel) handleHTTP(w http.ResponseWriter, req *http.Request) {
	req.URL.Scheme = t.targetURL.Scheme
	req.URL.Host = t.targetURL.Host
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		zap.L().Error("failed to dial upstream", zap.String("upstream", t.targetURL.Host))
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
