package handler

import (
	"API/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user repository.Users

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			responseJsonError(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var user_id int
		user_id, err = h.serv.CreateUser(&user)
		if err != nil {
			responseJsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var answerString string = "Account with username " + user.Username + " has created"
		responseJsonData(w, answerString, http.StatusOK)

		var address string = r.RemoteAddr
		log.Printf("Action: sign up from: %s, user_id = %d\n", address, user_id)

	default:
		responseJsonError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
