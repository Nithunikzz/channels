package main

import (
	"cop/channels/handlers"
	"cop/channels/models"
	"cop/channels/worker"
	"log"
	"net/http"
)

func main() {
	eventChan := make(chan models.InputEvent)

	// Start the worker
	go worker.StartWorker(eventChan)

	// Set up HTTP server and routes
	http.HandleFunc("/event", handlers.HandleEvent(eventChan))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
