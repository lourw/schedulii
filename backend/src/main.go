package main

import (
	"log"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatal("Unable to initialize app:", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
	defer app.Teardown()
}

