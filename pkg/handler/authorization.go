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
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		_, err = h.serv.CreateUser(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		var answerString string = "Account with username" + user.Username + "has created\n"
		w.Write([]byte(answerString))

		var address string = r.RemoteAddr
		log.Printf("Action: sign up from: %s\n", address)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
