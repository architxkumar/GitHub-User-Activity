package main

import (
	"GitHub-User-Activity/internal/helper"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid command usage: Invalid number of arguments")
		os.Exit(1)
	}
	username := arguments[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error connecting to GitHub\n", err.Error())
		os.Exit(1)
	}
	parsedResponseBody, err := helper.ParseResponseBody(response.Body)
	if err != nil {
		fmt.Println("Error parsing response body\n", err.Error())
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
		switch event.EventType {
		case "CommitCommentEvent":
			fmt.Println("Commented on commit on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "CreateEvent":
			fmt.Println("Created ", event.Payload.ReferenceType, "repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "DeleteEvent":
			fmt.Println("Deleted repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "ForkEvent":
			fmt.Println("Forked repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "GollumEvent":
			fmt.Println("Wiki update on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "IssueCommentEvent":
			fmt.Println("Issue comment on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "IssuesEvent":
			fmt.Println("Issue event on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "MemberEvent":
			fmt.Println("Member event on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "PublicEvent":
			fmt.Println(event.Repo.Name, " made public at ", event.CreatedAt)
		case "PullRequestEvent":
			fmt.Println("Pull request event on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "PullRequestReviewEvent":
			fmt.Println("Made a pull request review on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "PullRequestReviewCommentEvent":
			fmt.Println("Comment on a pull request review on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "PullRequestReviewThreadEvent":
			fmt.Println("Marked pull request review on repository ", event.Repo.Name, " as ", event.Payload.Action, " at ", event.CreatedAt)
		case "PushEvent":
			fmt.Println("Pushed ", event.Payload.Size, "commits to repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "ReleaseEvent":
			fmt.Println("Published package on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "SponsorshipEvent":
			fmt.Println("Created sponsorship listing on repository ", event.Repo.Name, " at ", event.CreatedAt)
		case "WatchEvent":
			fmt.Println("Starred repository ", event.Repo.Name, " at ", event.CreatedAt)
		default:
			fmt.Println("Unhandled event: ", event.EventType)
		}
	}
}
