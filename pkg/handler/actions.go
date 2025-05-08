package handler

import (
	_ "API/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	SenderName    string
	RecipientName string
	Text          string
}

func (h *Handler) message(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Action: get the message"))
	case http.MethodPost:
		message := Message{}

		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		err = h.serv.SendMessage(&message.SenderName, &message.RecipientName)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		var address string = r.RemoteAddr
		log.Printf("Action: sign up from: %s\n", address)

	case http.MethodPatch:
		w.Write([]byte("Action: edit the message"))
	case http.MethodDelete:
		w.Write([]byte("Action: delete the message"))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Action: get the message"))
	case http.MethodPost:
		w.Write([]byte("Action: send the message"))
	default:
		w.Write([]byte("This method is't allowed"))
	}
}
