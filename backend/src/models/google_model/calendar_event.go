package google_model

// Define a struct that represents a Google calendar event
type CalendarEvent struct {
	Summary string    `json:"summary"`
	Start   EventTime `json:"start"`
	End     EventTime `json:"end"`
}
