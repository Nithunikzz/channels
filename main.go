package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const webhookURL = "https://webhook.site/1139cb0e-d1ef-4c04-8e0e-01bc78532654"

// InputEvent represents the input event structure
type InputEvent struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	Id     string `json:"id"`
	Uid    string `json:"uid"`
	Mid    string `json:"mid"`
	T      string `json:"t"`
	P      string `json:"p"`
	L      string `json:"l"`
	Sc     string `json:"sc"`
	Atrk1  string `json:"atrk1"`
	Atrv1  string `json:"atrv1"`
	Atrt1  string `json:"atrt1"`
	Atrk2  string `json:"atrk2"`
	Atrv2  string `json:"atrv2"`
	Atrt2  string `json:"atrt2"`
	Atrk3  string `json:"atrk3"`
	Atrv3  string `json:"atrv3"`
	Atrt3  string `json:"atrt3"`
	Atrk4  string `json:"atrk4"`
	Atrv4  string `json:"atrv4"`
	Atrt4  string `json:"atrt4"`
	Uatrk1 string `json:"uatrk1"`
	Uatrv1 string `json:"uatrv1"`
	Uatrt1 string `json:"uatrt1"`
	Uatrk2 string `json:"uatrk2"`
	Uatrv2 string `json:"uatrv2"`
	Uatrt2 string `json:"uatrt2"`
	Uatrk3 string `json:"uatrk3"`
	Uatrv3 string `json:"uatrv3"`
	Uatrt3 string `json:"uatrt3"`
	Uatrk4 string `json:"uatrk4"`
	Uatrv4 string `json:"uatrv4"`
	Uatrt4 string `json:"uatrt4"`
	Uatrk5 string `json:"uatrk5"`
	Uatrv5 string `json:"uatrv5"`
	Uatrt5 string `json:"uatrt5"`
	Uatrk6 string `json:"uatrk6"`
	Uatrv6 string `json:"uatrv6"`
	Uatrt6 string `json:"uatrt6"`
}

// TransformedEvent represents the transformed event structure
type TransformedEvent struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	Traits          map[string]UserTrait `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type UserTrait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

var (
	wg sync.WaitGroup
)

func main() {
	eventChan := make(chan InputEvent)

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var event InputEvent
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eventChan <- event
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	go worker(eventChan)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func worker(eventChan chan InputEvent) {
	for event := range eventChan {
		wg.Add(1)
		go processEvent(event)
	}
}

func processEvent(event InputEvent) {
	defer wg.Done()

	attributes := make(map[string]Attribute)
	traits := make(map[string]UserTrait)

	// Handle attributes
	if event.Atrk1 != "" && event.Atrv1 != "" && event.Atrt1 != "" {
		attributes[event.Atrk1] = Attribute{Value: event.Atrv1, Type: event.Atrt1}
	}
	if event.Atrk2 != "" && event.Atrv2 != "" && event.Atrt2 != "" {
		attributes[event.Atrk2] = Attribute{Value: event.Atrv2, Type: event.Atrt2}
	}
	if event.Atrk3 != "" && event.Atrv3 != "" && event.Atrt3 != "" {
		attributes[event.Atrk3] = Attribute{Value: event.Atrv3, Type: event.Atrt3}
	}
	if event.Atrk4 != "" && event.Atrv4 != "" && event.Atrt4 != "" {
		attributes[event.Atrk4] = Attribute{Value: event.Atrv4, Type: event.Atrt4}
	}

	// Handle traits
	if event.Uatrk1 != "" && event.Uatrv1 != "" && event.Uatrt1 != "" {
		traits[event.Uatrk1] = UserTrait{Value: event.Uatrv1, Type: event.Uatrt1}
	}
	if event.Uatrk2 != "" && event.Uatrv2 != "" && event.Uatrt2 != "" {
		traits[event.Uatrk2] = UserTrait{Value: event.Uatrv2, Type: event.Uatrt2}
	}
	if event.Uatrk3 != "" && event.Uatrv3 != "" && event.Uatrt3 != "" {
		traits[event.Uatrk3] = UserTrait{Value: event.Uatrv3, Type: event.Uatrt3}
	}
	if event.Uatrk4 != "" && event.Uatrv4 != "" && event.Uatrt4 != "" {
		traits[event.Uatrk4] = UserTrait{Value: event.Uatrv4, Type: event.Uatrt4}
	}
	if event.Uatrk5 != "" && event.Uatrv5 != "" && event.Uatrt5 != "" {
		traits[event.Uatrk5] = UserTrait{Value: event.Uatrv5, Type: event.Uatrt5}
	}
	if event.Uatrk6 != "" && event.Uatrv6 != "" && event.Uatrt6 != "" {
		traits[event.Uatrk6] = UserTrait{Value: event.Uatrv6, Type: event.Uatrt6}
	}

	transformedEvent := TransformedEvent{
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
