package helper

import (
	"errors"
	"fmt"
	"net/http"
)

func MakeRequest(username string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	} else if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("unexpected status code: %d", response.StatusCode))
	}
	return response, nil
}
