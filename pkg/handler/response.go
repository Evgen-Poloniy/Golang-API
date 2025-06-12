package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func responseJsonMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	type jsonMessage struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	answer := jsonMessage{
		Code:    code,
		Message: message,
	}

	jsonData, err := json.MarshalIndent(answer, "", "    ")
	if err != nil {
		var errMessage string = fmt.Sprintf("Internal Server Error: %v\n", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func responseJsonError(w http.ResponseWriter, errorAnswer string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	type jsonError struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}

	answer := jsonError{
		Code:  code,
		Error: errorAnswer,
	}

	jsonData, err := json.MarshalIndent(answer, "", "    ")
	if err != nil {
		var errMessage string = fmt.Sprintf("Internal Server Error: %v\n", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func responseJsonData(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		var errMessage string = fmt.Sprintf("Internal Server Error: %v\n", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
