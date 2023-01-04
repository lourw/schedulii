package google

import (
	"context"
	"log"
	"net/http"
	models "schedulii/src/models"
	utils "schedulii/src/utils"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

// Gets a User's Calendar Events as a JSON file
func UserCalendarEventsHandler(c *gin.Context) {
	client, ok := utils.GetGoogleClient(c)
	if !ok {
		log.Fatalf("No google client or credentials found")
	}

	calendarEvents := getCalendarEvents(client)
	c.JSON(http.StatusOK, calendarEvents)
}

func getCalendarEvents(client *http.Client) []models.CalendarEvent {
	svc, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	today := time.Now().Format(time.RFC3339)
	furthestDateRetrieved := time.Now().AddDate(0, 1, 0).Format(time.RFC3339)

	// Retrieves the user's events
	events, err := svc.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(today).TimeMax(furthestDateRetrieved).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve the user's events for the next month: %v", err)
	}

	var calendarEvents []models.CalendarEvent
	for _, item := range events.Items {
		ce := models.CalendarEvent{
			Summary: item.Summary,
			Start: models.Time{
				DateTime: item.Start.DateTime,
				TimeZone: item.Start.TimeZone,
			},
			End: models.Time{
				DateTime: item.End.DateTime,
				TimeZone: item.End.TimeZone,
			},
		}
		calendarEvents = append(calendarEvents, ce)
	}

	return calendarEvents
}
