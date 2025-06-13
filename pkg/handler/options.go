package handler

import (
	"API/pkg/constant"
	"net/http"
	"os"
)

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var action string = "ping"
	var urlString string = r.URL.String()

	if r.Method == http.MethodGet {
		var statusCode int = http.StatusOK
		responseJsonMessage(w, "connection successful", statusCode)
		logEvent(address, action, urlString, r.Method, statusCode)

	} else {
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}

func (h *Handler) options(w http.ResponseWriter, r *http.Request) {
	var address string = r.RemoteAddr
	var action string = "option"
	var urlString string = r.URL.String()

	if r.Method == http.MethodGet {
		var statusCode int = http.StatusOK
		w.WriteHeader(statusCode)

		if os.Getenv("MODE") == constant.DEBUG_WITHOUT_DB {
			w.Write([]byte(h.printHandlers(urlsParametrsDebugWithoutDB)))
		} else {
			w.Write([]byte(h.printHandlers(urlsParametrs)))
		}

		logEvent(address, action, urlString, "GET", statusCode)

	} else {
		var statusCode int = http.StatusMethodNotAllowed
		responseJsonError(w, methodNotAllowed, statusCode)
		logError(address, action, urlString, r.Method, statusCode, methodNotAllowed)
	}
}
