package models

// Define a struct that represents a time in a Google calendar event
type Time struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}
