package models

import "time"

type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	EventType string    `json:"event_type"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEventRequest struct {
	UserID    int                    `json:"user_id"`
	EventType string                 `json:"event_type"`
	Metadata  map[string]interface{} `json:"metadata"`
}
