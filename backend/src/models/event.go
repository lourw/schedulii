package models

import "time"

type Event struct {
	EventId   int       `json:"eventId" form:"eventId" binding:"required"`
	GroupId   int       `json:"groupId" form:"groupId"`
	EventName string    `json:"eventName" form:"eventName"`
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"endTime" form:"endTime"`
}

