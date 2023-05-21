package middleware

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"go.uber.org/zap"
)

func HTTPLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := zap.L()
		ctx := context.New(r.Context()).Named("service").WithFields(
			zap.String("method", r.Method),
			zap.String("request_uri", r.RequestURI),
		).WithTrace()
		// Do stuff here
		logger.Debug("processing request",
			zap.String("request_uri", r.RequestURI))
		next.ServeHTTP(w, r.WithContext(ctx))
		logger.Debug("processed request")
	})
}

func HTTPErrorHandler(message string) http.Handler {
	logger := zap.L()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Debug(message,
			zap.String("request_uri", r.RequestURI))
	})
}
