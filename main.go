package main

import (
	"log"

	"loan/internal/bootstrap/app"
)

func main() {
	application, err := app.NewApp(".env")
	if err != nil {
		log.Panicf("Error initializing application: %v", err)
	}

	if err := application.Start(); err != nil {
		log.Panicf("Error starting application: %v", err)
	}
}
