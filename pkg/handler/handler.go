package handler

import (
	"API/pkg/constants"
	"API/pkg/service"
	"fmt"
	"net/http"
	"os"
)

var urlPaths []string

type Handler struct {
	serv *service.Service
}

func NewHendler(serv *service.Service) *Handler {
	return &Handler{serv}
}

func (h *Handler) printHandlers(urls *[]string) string {
	var lines string = "Handlers:\n"
	for _, str := range *urls {
		lines += "\t" + str + "\n"
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
		urls := []string{
			pathPing + " - GET",
			pathOption + " - GET",
			pathSingUp + " - POST",
			pathActionUser + " - GET",
			pathActionShutDownServer + " - GET",
		}
		urlPaths = urls

		router.HandleFunc(pathPing, h.ping)
		router.HandleFunc(pathOption, h.options)
		router.HandleFunc(pathSingUp, h.signUp)
		router.HandleFunc(pathActionUser, h.getUserById)
		router.HandleFunc(pathActionShutDownServer, h.shutdownServer)

		fmt.Print(h.printHandlers(&urlPaths))

	case constants.DEBUG_WITHOUT_DB:
		urls := []string{
			pathPing + " - GET",
			pathOption + " - GET",
			pathActionShutDownServer + " - GET",
		}
		urlPaths = urls

		router.HandleFunc(pathPing, h.ping)
		router.HandleFunc(pathOption, h.options)
		router.HandleFunc(pathActionShutDownServer, h.shutdownServer)

		fmt.Print(h.printHandlers(&urlPaths))
	}
}
