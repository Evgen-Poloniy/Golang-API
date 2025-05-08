package handler

import (
	//"API/pkg/repository"
	"API/pkg/service"
	"net/http"
	//"encoding/json"
)

type Handler struct {
	serv *service.Service
}

func NewHendler(serv *service.Service) *Handler {
	return &Handler{serv}
}

func (h *Handler) Handle() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/ping", h.ping)
	router.HandleFunc("/options", h.options)
	router.HandleFunc("/sign-up", h.signUp)
	router.HandleFunc("/sign-in", h.signIn)
	router.HandleFunc("/message", h.message)
	router.HandleFunc("/transaction", h.transaction)

	return router
}
