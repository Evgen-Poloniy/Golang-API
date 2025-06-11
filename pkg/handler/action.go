package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var user_id_str string = r.URL.Query().Get("user_id")
		if user_id_str == "" {
			responseJsonError(w, "Parameter user_id is required", http.StatusBadRequest)
			return
		}

		user_id, err := strconv.Atoi(user_id_str)
		if err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
		}

		user, err := h.serv.GetUserById(user_id)
		if err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJsonData(w, user, http.StatusOK)

		var address string = r.RemoteAddr
		log.Printf("Action: get information about user by user_id from: %s\n", address)

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getUserByUsername(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var username string = r.URL.Query().Get("username")
		if username == "" {
			responseJsonError(w, "Parameter username is required", http.StatusBadRequest)
			return
		}

		user, err := h.serv.GetUserByUsername(username)
		if err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJsonData(w, user, http.StatusOK)

		var address string = r.RemoteAddr
		log.Printf("Action: get information about user by username from: %s\n", address)

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getUserByAttributes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var fields allowedAttributes

		err := json.NewDecoder(r.Body).Decode(&fields)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			responseJsonError(w, "Values of attributus is required", http.StatusBadRequest)
			return
		}

		user, err := h.serv.GetUserByAttributes(attributes)
		if err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJsonData(w, user, http.StatusOK)

		var address string = r.RemoteAddr
		log.Printf("Action: get information about user from: %s\n", address)

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) shutdownServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var password string = r.URL.Query().Get("password")
		if password == "" {
			responseJsonError(w, "Parameter password is required", http.StatusBadRequest)
			return
		}

		if password == os.Getenv("API_PASSWORD") {
			responseJsonData(w, "Server is shutting down", http.StatusOK)
		} else {
			responseJsonError(w, "Whrong password. Try again", http.StatusBadRequest)
		}

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) testGetUserByAttributes(w http.ResponseWriter, r *http.Request) {
	var field allowedAttributes

	err := json.NewDecoder(r.Body).Decode(&field)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(field)

	attributes := make(map[string]string, countAllowedAttributes)

	if field.ID != 0 {
		attributes["user_id"] = strconv.FormatUint(uint64(field.ID), 10)
	}
	if field.Username != "" {
		attributes["username"] = field.Username
	}
	if field.Name != "" {
		attributes["name"] = field.Name
	}
	if field.Surname != "" {
		attributes["surname"] = field.Surname
	}

	for key, value := range attributes {
		fmt.Println(key, value)
	}

}
