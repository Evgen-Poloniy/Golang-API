package handler

import (
	"log"
)

func logError(address string, action string, URL, method string, statusCode int, errorLog string) {
	log.Printf("Request from address: %s, action: %s, URL: %s, method: %s, status: %d,\nError: %s\n", address, action, URL, method, statusCode, errorLog)
}

func logEvent(address string, action string, URL string, method string, statusCode int) {
	log.Printf("Request from address: %s, action: %s, URL: %s, method: %s, status: %d\n", address, action, URL, method, statusCode)
}
