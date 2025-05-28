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

	router.HandleFunc(pathPing, h.ping)
	router.HandleFunc(pathOptions, h.options)
	router.HandleFunc(pathSingUp, h.signUp)

	return router
}
