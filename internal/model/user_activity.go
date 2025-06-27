package model

import "time"

// UserActivity represents the cached user entry in local file
type UserActivity struct {
	Username string `json:"username"`
	// Represents the list of user Event from GitHub Events API
	Content []Event `json:"content"`
	// Timestamp is used when comparing the cache expiry
	Timestamp time.Time `json:"timestamp"`
}
