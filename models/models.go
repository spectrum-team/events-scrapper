package models

import (
	"time"
)

// Event represent a single event happening in Santo Domingo
type Event struct {
	Name      string    `json:"name,omitempty"`
	EventType string    `json:"eventtype,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Place     string    `json:"place,omitempty"`
}
