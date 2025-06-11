package handler

import (
	"API/pkg/constants"
	"log"
	"net/http"
	"os"
)

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	responseJsonMessage(w, "Connection successful", http.StatusOK)

	var address string = r.RemoteAddr
	log.Printf("Checking connection from address: %s", address)
}

func (h *Handler) options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if os.Getenv("MODE") == constants.DEBUG_WITHOUT_DB {
		w.Write([]byte(h.printHandlers(urlsParametrsDebugWithoutDB)))
	} else {
		w.Write([]byte(h.printHandlers(urlsParametrs)))
	}

	var address string = r.RemoteAddr
	log.Printf("Requested actions of the server from address: %s", address)
}
