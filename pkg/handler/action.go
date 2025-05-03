package handler

import (
	"net/http"
)

func (h *Handler) message(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Action: get the message"))
	case http.MethodPost:
		w.Write([]byte("Action: send the message"))
	case http.MethodPut:
		w.Write([]byte("Action: edit the message"))
	case http.MethodDelete:
		w.Write([]byte("Action: delete the message"))
	default:
		w.Write([]byte("This method is't allowed"))
	}
}

func (h *Handler) transaction(w http.ResponseWriter, r *http.Request) {

}
