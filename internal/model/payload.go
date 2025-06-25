package model

// The event payload object is unique to the event type.
// It encompasses necessary properties from all different event types.
type payload struct {
	// The type of Git ref object created in the repository.
	// Can be either branch, tag, or repository.
	//
	// Specific to CreateEvent.
	//
	// For more information, visit: https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#event-payload-object-for-createevent
	ReferenceType string `json:"ref_type"`
	// The number of commits in the push.
	//
	// Specific to PushEvent.
	//
	// For more information, visit: https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#event-payload-object-for-pushevent
	Size int `json:"size"`
	// The specific action performed in the event
	//
	// Unique to each event type.
	//
	// For more information, visit: https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#event-object-common-properties
	Action string `json:"action"`
}
