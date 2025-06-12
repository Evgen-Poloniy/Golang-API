package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type transactionData struct {
	UsernameSender    string  `json:"username_sender"`
	UsernameRecipient string  `json:"username_recipient"`
	Amount            float64 `json:"amount"`
}

func (t *transactionData) Validate() error {
	var missingFields []string

	if t.UsernameSender == "" {
		missingFields = append(missingFields, "username_sender")
	}
	if t.UsernameRecipient == "" {
		missingFields = append(missingFields, "username_recipient")
	}
	if t.Amount == 0 {
		missingFields = append(missingFields, "amount")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required fields: %v", strings.Join(missingFields, ", "))
	}
	return nil
}

func (h *Handler) makeTransaction(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "makeTransaction"

	switch r.Method {
	case http.MethodPatch:

		var data transactionData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			var statusCode int = http.StatusBadRequest
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if err := data.Validate(); err != nil {
			var statusCode int = http.StatusBadRequest
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if data.Amount < 0 {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "amount can't be less then 0"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if data.UsernameSender == data.UsernameRecipient {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "the sender and the recipient are one user. Operation denied"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if err := h.serv.MakeTransaction(data.UsernameSender, data.UsernameRecipient, data.Amount, h.serv.Action); err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonMessage(w, "transaction was successful", statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
