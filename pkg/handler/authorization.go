package handler

import (
	"API/pkg/attribute"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func validate(auth *attribute.AuthField) error {
	var missingFields []string

	if auth.Username == "" {
		missingFields = append(missingFields, "username")
	}
	if auth.Password == "" {
		missingFields = append(missingFields, "password")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required fields: %v", strings.Join(missingFields, ", "))
	}

	return nil
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "signUp"

	switch r.Method {
	case http.MethodPost:
		var auth attribute.AuthField

		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "invalid input"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if err := validate(&auth); err != nil {
			var statusCode int = http.StatusBadRequest
			var errorStr string = err.Error()
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		_, err = h.serv.CreateUser(&auth)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		var answerString string = "account with username '" + auth.Username + "' has created"
		responseJsonMessage(w, answerString, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
