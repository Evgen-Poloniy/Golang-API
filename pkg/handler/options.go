package handler

import (
	serverHTTP "API/pkg/server"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	var resp serverHTTP.Response = serverHTTP.Response{
		Status:  http.StatusOK,
		Message: "Connection successful",
	}

	serverHTTP.JsonResponce(resp, w)

	var address string = r.RemoteAddr
	log.Printf("Checking connection from address: %s", address)
}

func (h *Handler) options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	lines := []string{
		"/ping - GET",
		"/options - GET",
		"/sign-up - POST",
		"/sign-in - POST",
		"/message - GET, POST, PATCH, DELETE",
		"/transaction - GET, POST",
	}
	w.Write([]byte(strings.Join(lines, "\n")))

	var address string = r.RemoteAddr
	log.Printf("Requested actions of the server from address: %s", address)
}
