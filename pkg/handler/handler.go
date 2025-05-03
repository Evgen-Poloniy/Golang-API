package handler

import (
	//"API/pkg/repository"
	serverHTTP "API/pkg/server"
	"API/pkg/service"
	"net/http"
	"strings"

	//"encoding/json"
	"log"
)

type Handler struct {
	srv *service.Service
}

func NewHendler(srv *service.Service) *Handler {
	return &Handler{srv}
}

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

func (h *Handler) Handle() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/ping", h.ping)
	router.HandleFunc("/options", h.options)
	router.HandleFunc("/sing-up", h.signUp)
	router.HandleFunc("/sing-in", h.signIn)
	router.HandleFunc("/message", h.message)
	router.HandleFunc("/transaction", h.transaction)

	return router
}

/*
func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	var user repository.Users
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return err
    }

	if err := h.srv.Authorization.CreateAccount(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return nil
}
*/
