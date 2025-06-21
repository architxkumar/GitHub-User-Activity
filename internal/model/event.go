package model

import "time"

type Event struct {
	Id        string     `json:"id"`
	EventType string     `json:"type"`
	Actor     Actor      `json:"actor"`
	Repo      repository `json:"repo"`
	Public    bool       `json:"public"`
	CreatedAt time.Time  `json:"created_at"`
	Payload   payload    `json:"payload"`
}
