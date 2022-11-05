package service

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	if len(headers) > 0 {
		for _, headerGroup := range headers {
			for key, values := range headerGroup {
				w.Header()[key] = values
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// errorJSON takes an error and generates and sends a json error response with a BadRequest status code.
func errorJSON(w http.ResponseWriter, err error) error {
	return errorJSONWithStatus(w, err, http.StatusBadRequest)
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func errorJSONWithStatus(w http.ResponseWriter, err error, statusCode int) error {
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return writeJSON(w, statusCode, payload)
}
