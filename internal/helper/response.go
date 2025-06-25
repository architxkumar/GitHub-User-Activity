package helper

import (
	"GitHub-User-Activity/internal/model"
	"encoding/json"
	"io"
)

// ParseResponseBody parses an HTTP response body into a slice of model.Event.
// It reads all bytes from the response body and unmarshalls the JSON data.
// Returns the parsed events on success, or an error if reading or unmarshalling fails.
func ParseResponseBody(responseBody io.ReadCloser) ([]model.Event, error) {
	body, err := io.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}
	var eventList []model.Event
	err = json.Unmarshal(body, &eventList)
	if err != nil {
		return nil, err
	}
	return eventList, nil
}
