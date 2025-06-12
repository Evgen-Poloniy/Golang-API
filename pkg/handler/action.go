package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "getUserByID"

	switch r.Method {
	case http.MethodGet:
		var user_id_str string = r.URL.Query().Get("user_id")
		if user_id_str == "" {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "Parameter user_id is required"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		user_id, err := strconv.Atoi(user_id_str)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		user, err := h.serv.GetUserByID(user_id)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonData(w, user, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}

func (h *Handler) getUserByUsername(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "getUserByUsername"

	switch r.Method {
	case http.MethodGet:
		var username string = r.URL.Query().Get("username")
		if username == "" {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "parameter username is required"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		user, err := h.serv.GetUserByUsername(username)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonData(w, user, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}

func (h *Handler) getUserByAttributes(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "getUserByAttributes"

	switch r.Method {
	case http.MethodPost:
		var fields allowedAttributes

		err := json.NewDecoder(r.Body).Decode(&fields)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		attributes := make(map[string]string, countAllowedAttributes)

		if fields.ID != 0 {
			attributes["user_id"] = strconv.FormatUint(uint64(fields.ID), 10)
		}
		if fields.Username != "" {
			attributes["username"] = fields.Username
		}
		if fields.Name != "" {
			attributes["name"] = fields.Name
		}
		if fields.Surname != "" {
			attributes["surname"] = fields.Surname
		}

		if len(attributes) == 0 {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "values of attributus is required"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		user, err := h.serv.GetUserByAttributes(attributes)
		if err != nil {
			var statusCode int = http.StatusInternalServerError
			var errorStr string = err.Error()
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonData(w, user, statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}

func (h *Handler) shutdownServer(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var urlString string = r.URL.String()
	var action string = "shutdownServer"

	switch r.Method {
	case http.MethodGet:
		var password string = r.URL.Query().Get("password")
		if password == "" {
			var statusCode int = http.StatusBadRequest
			var errorStr string = "parameter password is required"
			responseJsonError(w, errorStr, statusCode)
			logError(address, action, urlString, r.Method, statusCode, errorStr)
			return
		}

		if password == os.Getenv("API_PASSWORD") {
			var statusCode int = http.StatusOK
			responseJsonData(w, "server is shutting down", statusCode)
		} else {
			var statusCode int = http.StatusBadRequest
			responseJsonError(w, "whrong password. Try again", statusCode)
			return
		}

		var statusCode int = http.StatusOK
		responseJsonMessage(w, "server has shuted down", statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	default:
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
