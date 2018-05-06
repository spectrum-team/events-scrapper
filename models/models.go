package models

import (
	"time"
)

// Event represent a single event happening in Santo Domingo
type Event struct {
	Name      string    `json:"name"`
	EventType string    `json:"type"`
	Date      time.Time `json:"date"`
	Place     string    `json:"place"`
}
