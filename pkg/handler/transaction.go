package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
	switch r.Method {
	case http.MethodPatch:

		var data transactionData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			responseJsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := data.Validate(); err != nil {
			responseJsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if data.Amount < 0 {
			responseJsonError(w, "amount can't be less then 0", http.StatusBadRequest)
			return
		}

		if data.UsernameSender == data.UsernameRecipient {
			responseJsonError(w, "the sender and the recipient are one user. Operation denied", http.StatusBadRequest)
			return
		}

		if err := h.serv.MakeTransaction(data.UsernameSender, data.UsernameRecipient, data.Amount, h.serv.Action); err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var address string = r.RemoteAddr
		log.Printf("Action: get information about user from: %s\n", address)

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
