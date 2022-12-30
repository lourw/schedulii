package models

// Define a struct that represents a Google calendar
type CalendarEntry struct {
	Id       string `json:"id"`
	Location     string `json:"location"`
	Summary  string `json:"summary"`
	TimeZone string `json:"timeZone"`
}
