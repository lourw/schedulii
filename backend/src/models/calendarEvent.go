package models

// Define a struct that represents a Google calendar event
type CalendarEvent struct {
	Summary string `json:"summary"`
	Start   Time   `json:"start"`
	End     Time   `json:"end"`
}
