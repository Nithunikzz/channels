package models

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
