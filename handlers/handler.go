package handlers

import (
	"cop/channels/models"
	"encoding/json"
	"net/http"
)

// HandleEvent handles incoming event requests and sends them to the event channel.
func HandleEvent(eventChan chan<- models.InputEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var event models.InputEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eventChan <- event
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}
}
