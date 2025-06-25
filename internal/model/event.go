package model

import "time"

// Event represents the response from the events API endpoint
type Event struct {
	// Id is the unique identifier for the event.
	Id string `json:"id"`
	// EventType is the type of event. Events uses PascalCase for the name.
	EventType string `json:"type"`
	// Actor is the user that triggered the event.
	Actor Actor `json:"actor"`
	// Repo is the repository object where the event occurred.
	Repo repository `json:"repo"`
	// Public represents whether the event is visible to all users.
	Public bool `json:"public"`
	// CreatedAt is the date and time when the event was triggered. It is formatted according to ISO 8601.
	CreatedAt time.Time `json:"created_at"`
	// Payload is the event payload object is unique to the event type.
	Payload payload `json:"payload"`
}
