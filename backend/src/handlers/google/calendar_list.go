package google

import (
	"context"
	"log"
	"net/http"
	models "schedulii/src/models"
	utils "schedulii/src/utils"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

func UserCalendarListHandler(c *gin.Context) {
	client, ok := utils.GetGoogleClient(c)
	if !ok {
		log.Fatalf("No google client or credentials found")
	}

	calendars := getUserCalendarList(client)
	c.JSON(http.StatusOK, calendars)
}

func getUserCalendarList(client *http.Client) []models.Calendar {
	svc, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	calendarList, err := svc.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve the user's list of calendars: %v", err)
	}

	var calendars []models.Calendar
	for _, item := range calendarList.Items {
		ce := models.Calendar{
			Id:       item.Id,
			Location: item.Location,
			Summary:  item.Summary,
			TimeZone: item.TimeZone,
		}
		calendars = append(calendars, ce)
	}

	return calendars
}
