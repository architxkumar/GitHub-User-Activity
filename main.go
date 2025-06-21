package main

import (
	"GitHub-User-Activity/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid number of arguments")
	}
	username := arguments[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to connect to Github.\n", err.Error())
	}
	switch response.StatusCode {
	case 200:
		printPublicUserEvents(response.Body)
	case 404:
		fmt.Println("User not found")
	default:
		fmt.Println(response.StatusCode, response.Status)
	}
	if err := response.Body.Close(); err != nil {
		log.Fatal("Unable to close response body.\n", err.Error())
	}
}

func parseResponseBody(responseBody io.ReadCloser) []model.Event {
	body, err := io.ReadAll(responseBody)
	if err != nil {
		log.Fatal("Unable to read response body.\n", err.Error())
	}
	var eventArray []model.Event
	err = json.Unmarshal(body, &eventArray)
	if err != nil {
		log.Fatal("Unable to parse JSON.\n", err.Error())
	}
	return eventArray
}

func printPublicUserEvents(responseBody io.ReadCloser) {
	eventArray := parseResponseBody(responseBody)
	if len(eventArray) == 0 {
		fmt.Println("No events found")
	}
	for _, event := range eventArray {
		fmt.Println(event.EventType, " to ", event.Repo.Name, " at ", event.CreatedAt)
	}
}
