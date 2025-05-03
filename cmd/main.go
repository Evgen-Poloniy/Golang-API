package main

import (
	//"database/sql"

	"API/pkg/handler"
	"API/pkg/repository"
	serverHTTP "API/pkg/server"
	"API/pkg/service"

	"log"
	"time"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	hand := handler.NewHendler(serv)

	srv := serverHTTP.NewServer("localhost", "3505", 1<<20, hand.Handle(), 10*time.Second, 10*time.Second)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
