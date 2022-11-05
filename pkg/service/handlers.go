package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/models"
)

type authenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authenticateResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Authentication) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload authenticateRequest

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// validate the user against the database
	user, err := models.UserByEmail(r.Context(), app.conn, requestPayload.Email)
	if err != nil {
		errorJSONWithStatus(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		errorJSONWithStatus(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	payload := authenticateResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	writeJSON(w, http.StatusAccepted, payload)
}
