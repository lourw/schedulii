package google_model

// Define a struct that represents a time in a Google calendar event
type EventTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}
