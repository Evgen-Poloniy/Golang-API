package handler

import (
	"API/pkg/constants"
	"API/pkg/service"
	"fmt"
	"net/http"
	"os"
)

type Handler struct {
	serv *service.Service
}

func NewHendler(serv *service.Service) *Handler {
	return &Handler{serv}
}

func (h *Handler) printHandlers(urlsParametrs []handlerParametrs) string {
	var lines string = "Handlers:\n"
	for _, uprs := range urlsParametrs {
		lines += fmt.Sprintf("\t%s - %s\n", uprs.Path, uprs.Methods)
	}

	return lines
}

func (h *Handler) Handle() http.Handler {
	router := http.NewServeMux()

	checkMode(h, router)

	return router
}

func checkMode(h *Handler, router *http.ServeMux) {
	var mode string = os.Getenv("MODE")

	switch mode {
	case constants.DEBUG:
		router.HandleFunc(pathPing, h.ping)
		router.HandleFunc(pathOption, h.options)
		router.HandleFunc(pathSingUp, h.signUp)
		router.HandleFunc(pathActionUserSearch, h.getUserByAttributes)
		router.HandleFunc(pathActionUserGetByID, h.getUserByID)
		router.HandleFunc(pathActionUserGetByUsername, h.getUserByUsername)
		router.HandleFunc(pathActionShutDownServer, h.shutdownServer)
		router.HandleFunc(pathMakeTransaction, h.makeTransaction)
		router.HandleFunc(pathTransactionGetByID, h.getTransactionByID)

		fmt.Print(h.printHandlers(urlsParametrs))

	case constants.DEBUG_WITHOUT_DB:
		router.HandleFunc(pathPing, h.ping)
		router.HandleFunc(pathOption, h.options)
		router.HandleFunc(pathActionShutDownServer, h.shutdownServer)

		fmt.Print(h.printHandlers(urlsParametrsDebugWithoutDB))
	}
}
