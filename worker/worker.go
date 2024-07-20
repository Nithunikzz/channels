package worker

import (
	"bytes"
	"cop/channels/models"
	"encoding/json"
	"log"
	"net/http"
)

const webhookURL = "https://webhook.site/1139cb0e-d1ef-4c04-8e0e-01bc78532654"

// StartWorker starts a worker that processes events from the event channel.
func StartWorker(eventChan <-chan models.InputEvent) {
	for event := range eventChan {
		processEvent(event)
	}
}

func processEvent(event models.InputEvent) {
	attributes := make(map[string]models.Attribute)
	traits := make(map[string]models.UserTrait)

	// Handle attributes
	if event.Atrk1 != "" && event.Atrv1 != "" && event.Atrt1 != "" {
		attributes[event.Atrk1] = models.Attribute{Value: event.Atrv1, Type: event.Atrt1}
	}
	if event.Atrk2 != "" && event.Atrv2 != "" && event.Atrt2 != "" {
		attributes[event.Atrk2] = models.Attribute{Value: event.Atrv2, Type: event.Atrt2}
	}
	if event.Atrk3 != "" && event.Atrv3 != "" && event.Atrt3 != "" {
		attributes[event.Atrk3] = models.Attribute{Value: event.Atrv3, Type: event.Atrt3}
	}
	if event.Atrk4 != "" && event.Atrv4 != "" && event.Atrt4 != "" {
		attributes[event.Atrk4] = models.Attribute{Value: event.Atrv4, Type: event.Atrt4}
	}

	// Handle traits
	if event.Uatrk1 != "" && event.Uatrv1 != "" && event.Uatrt1 != "" {
		traits[event.Uatrk1] = models.UserTrait{Value: event.Uatrv1, Type: event.Uatrt1}
	}
	if event.Uatrk2 != "" && event.Uatrv2 != "" && event.Uatrt2 != "" {
		traits[event.Uatrk2] = models.UserTrait{Value: event.Uatrv2, Type: event.Uatrt2}
	}
	if event.Uatrk3 != "" && event.Uatrv3 != "" && event.Uatrt3 != "" {
		traits[event.Uatrk3] = models.UserTrait{Value: event.Uatrv3, Type: event.Uatrt3}
	}
	if event.Uatrk4 != "" && event.Uatrv4 != "" && event.Uatrt4 != "" {
		traits[event.Uatrk4] = models.UserTrait{Value: event.Uatrv4, Type: event.Uatrt4}
	}
	if event.Uatrk5 != "" && event.Uatrv5 != "" && event.Uatrt5 != "" {
		traits[event.Uatrk5] = models.UserTrait{Value: event.Uatrv5, Type: event.Uatrt5}
	}
	if event.Uatrk6 != "" && event.Uatrv6 != "" && event.Uatrt6 != "" {
		traits[event.Uatrk6] = models.UserTrait{Value: event.Uatrv6, Type: event.Uatrt6}
	}

	transformedEvent := models.TransformedEvent{
		Event:           event.Ev,
		EventType:       event.Et,
		AppID:           event.Id,
		UserID:          event.Uid,
		MessageID:       event.Mid,
		PageTitle:       event.T,
		PageURL:         event.P,
		BrowserLanguage: event.L,
		ScreenSize:      event.Sc,
		Attributes:      attributes,
		Traits:          traits,
	}

	transformedData, err := json.Marshal(transformedEvent)
	if err != nil {
		log.Println("Error marshalling transformed event:", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(transformedData))
	if err != nil {
		log.Println("Error sending transformed event to webhook:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Sent transformed event to webhook, status:", resp.Status)
}
