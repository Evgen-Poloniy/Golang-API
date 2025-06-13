package handler

import (
	"API/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

		senderID, recipientID, errTransaction := h.serv.MakeTransaction(data.UsernameSender, data.UsernameRecipient, data.Amount)
		if errTransaction != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = errTransaction.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		transaction := repository.Transactions{
			Amount:      data.Amount,
			SenderID:    senderID,
			RecipientID: recipientID,
		}

		_, err := h.serv.CreateRecordOfTransaction(&transaction)
		if err != nil {
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

func (h *Handler) getTransactionByID(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var action string = "getTransactionByID"
	var urlString string = r.URL.String()

	if r.Method == http.MethodGet {
		transaction_id_str := r.URL.Query().Get("transaction_id")
		if transaction_id_str == "" {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = "parameter 'transaction_id' is required"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		transaction_id_64, err := strconv.ParseUint(transaction_id_str, 10, 32)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		transaction_id := uint32(transaction_id_64)

		transaction, err := h.serv.GetTransactionByID(transaction_id)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonData(w, transaction, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	} else {
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
