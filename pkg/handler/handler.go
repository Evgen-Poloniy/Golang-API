package handler

import (
	//"API/pkg/repository"
	"API/pkg/service"
	"fmt"
	"net/http"
	//"encoding/json"
)

type Handler struct {
	serv *service.Service
}

func NewHendler(serv *service.Service) *Handler {
	return &Handler{serv}
}

func (h *Handler) printHandlers() string {
	lines := fmt.Sprint(
		"Handlers:\n",
		"\t"+pathPing+" - GET\n",
		"\t"+pathOption+" - GET\n",
		"\t"+pathSingUp+" - POST\n",
		"\t"+pathActionUser+" - GET\n",
	)

	return lines
}

func (h *Handler) Handle() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(pathPing, h.ping)
	router.HandleFunc(pathOption, h.options)
	router.HandleFunc(pathSingUp, h.signUp)
	router.HandleFunc(pathActionUser, h.getUserById)
	router.HandleFunc(pathActionShutDownServer, h.shutdownServer)

	fmt.Print(h.printHandlers())

	return router
}
