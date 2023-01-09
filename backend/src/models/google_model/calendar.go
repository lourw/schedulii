package google_model

// Define a struct that represents a Google calendar
type Calendar struct {
	Id       string `json:"id"`
	Location string `json:"location"`
	Summary  string `json:"summary"`
	TimeZone string `json:"timeZone"`
}
