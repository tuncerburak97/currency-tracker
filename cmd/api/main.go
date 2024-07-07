package main

import (
	"currency-tracker/internal/app"
	"log"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Could not initialize app: %v", err)
	}

	application.Start()
}
