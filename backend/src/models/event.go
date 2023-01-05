package models

type Event struct {
	EventName string `json:"title" binding:"required"`
	EventId   int    `json:"id" binding:"required"`
}
