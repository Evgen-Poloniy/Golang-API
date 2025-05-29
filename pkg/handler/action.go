package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var user_id_str string = r.URL.Query().Get("user_id")
		if user_id_str == "" {
			http.Error(w, "Parameter user_id is required", http.StatusBadRequest)
			return
		}

		user_id, convertationErr := strconv.Atoi(user_id_str)
		if convertationErr != nil {
			http.Error(w, "Bad convertation user_id to integer value", http.StatusInternalServerError)
		}

		user, err := h.serv.GetUserById(user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		var answer string = fmt.Sprintf(
			"{\n\t%d,\n\t\"%s\",\n\t\"%s\",\n\t\"%s\",\n\t\"%s\",\n\t%.2f\n}",
			user.ID,
			user.Username,
			user.Name,
			user.Surname,
			user.Password,
			user.Coins,
		)

		w.Write([]byte(answer))

		var address string = r.RemoteAddr
		log.Printf("Action: get information about user by user_id from: %s\n", address)

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) shutdownServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var password string = r.URL.Query().Get("password")
		if password == "" {
			http.Error(w, "Parameter password is required", http.StatusBadRequest)
			return
		}

		if password == os.Getenv("API_PASSWORD") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Server is shutting down..."))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "Parameter password is required", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
