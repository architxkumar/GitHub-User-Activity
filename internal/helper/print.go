package helper

import (
	"GitHub-User-Activity/internal/model"
	"fmt"
)

// PrintPublicUserEvents prints the public user events fetched from the GitHub Events API.
func PrintPublicUserEvents(eventList []model.Event) {
	if len(eventList) == 0 {
		fmt.Println("No events found")
	}
	for _, event := range eventList {
		message := formatEventMessage(event)
		if message != "" {
			fmt.Println(message)
		}
	}
}

// formatEventMessage returns formatted message for each event.
// In case, the argument is invalid, it returns an empty string
func formatEventMessage(event model.Event) string {
	// Checking for existence of guaranteed fields in the API response
	if event.EventType == "" || event.Id == "" || event.Repo.Name == "" || event.CreatedAt.IsZero() {
		return ""
	}
	var message string
	switch event.EventType {
	case "CommitCommentEvent":
		message = fmt.Sprintln("Commented on a commit on repository", event.Repo.Name, "at", event.CreatedAt)
	case "CreateEvent":
		message = fmt.Sprintln("Created", event.Payload.ReferenceType, event.Repo.Name, "at", event.CreatedAt)
	case "DeleteEvent":
		message = fmt.Sprintln("Deleted repository", event.Repo.Name, "at", event.CreatedAt)
	case "ForkEvent":
		message = fmt.Sprintln("Forked repository", event.Repo.Name, "at", event.CreatedAt)
	case "GollumEvent":
		message = fmt.Sprintln("Wiki update on repository", event.Repo.Name, "at", event.CreatedAt)
	case "IssueCommentEvent":
		message = fmt.Sprintln("Issue comment on repository", event.Repo.Name, "at", event.CreatedAt)
	case "IssuesEvent":
		message = fmt.Sprintln("Issue event on repository", event.Repo.Name, "at", event.CreatedAt)
	case "MemberEvent":
		message = fmt.Sprintln("Member event on repository", event.Repo.Name, "at", event.CreatedAt)
	case "PublicEvent":
		message = fmt.Sprintln(event.Repo.Name, "made public at", event.CreatedAt)
	case "PullRequestEvent":
		message = fmt.Sprintln("Pull request on repository", event.Repo.Name, "at", event.CreatedAt)
	case "PullRequestReviewEvent":
		message = fmt.Sprintln("Made a pull request review on repository", event.Repo.Name, "at", event.CreatedAt)
	case "PullRequestReviewCommentEvent":
		message = fmt.Sprintln("Comment on a pull request review on repository", event.Repo.Name, "at", event.CreatedAt)
	case "PullRequestReviewThreadEvent":
		message = fmt.Sprintln("Marked pull request review on repository", event.Repo.Name, "as", event.Payload.Action, "at", event.CreatedAt)
	case "PushEvent":
		// Conditional used to flip between singular and multiple commits while printing
		if event.Payload.Size == 1 {
			message = fmt.Sprintln("Pushed 1 commit to repository ", event.Repo.Name, "at", event.CreatedAt)
		} else {
			message = fmt.Sprintln("Pushed", event.Payload.Size, "commits to repository ", event.Repo.Name, " at ", event.CreatedAt)
		}
	case "ReleaseEvent":
		message = fmt.Sprintln("Published package on repository", event.Repo.Name, "at", event.CreatedAt)
	case "SponsorshipEvent":
		message = fmt.Sprintln("Created sponsorship listing on repository", event.Repo.Name, "at", event.CreatedAt)
	case "WatchEvent":
		message = fmt.Sprintln("Starred repository", event.Repo.Name, "at", event.CreatedAt)
	default:
		message = fmt.Sprintln("Performed event", event.EventType, "on", event.Repo.Name, "at", event.CreatedAt)
	}
	return message
}
