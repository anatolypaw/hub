package main

import (
	"hub/internal/app"
	"hub/internal/domain/models"
	"log"
	"os"
)

func main() {
	log.Print("Version ", models.Version)
	app.Run()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	}

}
