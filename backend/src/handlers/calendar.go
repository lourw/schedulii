package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

// Define a struct that represents a Google calendar event
type calendarEvent struct {
	Summary string `json:"summary"`
	Start   Time   `json:"start"`
	End     Time   `json:"end"`
}

// Define a struct that represents a time in a Google calendar event
type Time struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

func GetCalendars(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// Gets a User's Calendar Events as a JSON file
func GetUserCalendarEvents(c *gin.Context) {
	ctx := context.Background()

	config := readGoogleAPICredentials()

	tok, valid := tokenFromSession(c)
	if !valid {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		c.Redirect(http.StatusFound, authURL)
		c.Abort()
		return
	}

	client := config.Client(ctx, tok)

	// Creates a Google Calendar Service
	svc, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	today := time.Now().Format(time.RFC3339)
	one_month := time.Now().AddDate(0, 1, 0).Format(time.RFC3339)

	// Retrieves the user's events
	events, err := svc.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(today).TimeMax(one_month).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve the user's events for the next month: %v", err)
	}

	var calendarEvents []calendarEvent
	for _, item := range events.Items {
		ce := calendarEvent{
			Summary: item.Summary,
			Start: Time{
				DateTime: item.Start.DateTime,
				TimeZone: item.Start.TimeZone,
			},
			End: Time{
				DateTime: item.End.DateTime,
				TimeZone: item.End.TimeZone,
			},
		}
		calendarEvents = append(calendarEvents, ce)
	}

	b, err := json.Marshal(calendarEvents)
	if err != nil {
		log.Fatalf("Unable to marshal events: %v", err)
	}

	c.JSON(http.StatusOK, b)
}
