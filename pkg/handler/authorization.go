package handler

import (
	"API/pkg/repository"
	"encoding/json"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "signUp"

	switch r.Method {
	case http.MethodPost:
		var user repository.Users

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "invalid input"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		_, err = h.serv.CreateUser(&user)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		var answerString string = "account with username '" + user.Username + "' has created"
		responseJsonMessage(w, answerString, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
