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
		switch event.EventType {
		case "CommitCommentEvent":
			fmt.Println("Commented on a commit on repository", event.Repo.Name, "at", event.CreatedAt)
		case "CreateEvent":
			fmt.Println("Created", event.Payload.ReferenceType, event.Repo.Name, "at", event.CreatedAt)
		case "DeleteEvent":
			fmt.Println("Deleted repository", event.Repo.Name, "at", event.CreatedAt)
		case "ForkEvent":
			fmt.Println("Forked repository", event.Repo.Name, "at", event.CreatedAt)
		case "GollumEvent":
			fmt.Println("Wiki update on repository", event.Repo.Name, "at", event.CreatedAt)
		case "IssueCommentEvent":
			fmt.Println("Issue comment on repository", event.Repo.Name, "at", event.CreatedAt)
		case "IssuesEvent":
			fmt.Println("Issue event on repository", event.Repo.Name, "at", event.CreatedAt)
		case "MemberEvent":
			fmt.Println("Member event on repository", event.Repo.Name, "at", event.CreatedAt)
		case "PublicEvent":
			fmt.Println(event.Repo.Name, "made public at", event.CreatedAt)
		case "PullRequestEvent":
			fmt.Println("Pull request on repository", event.Repo.Name, "at", event.CreatedAt)
		case "PullRequestReviewEvent":
			fmt.Println("Made a pull request review on repository", event.Repo.Name, "at", event.CreatedAt)
		case "PullRequestReviewCommentEvent":
			fmt.Println("Comment on a pull request review on repository", event.Repo.Name, "at", event.CreatedAt)
		case "PullRequestReviewThreadEvent":
			fmt.Println("Marked pull request review on repository", event.Repo.Name, "as", event.Payload.Action, "at", event.CreatedAt)
		case "PushEvent":
			// Conditional used to flip between singular and multiple commits while printing
			if event.Payload.Size == 1 {
				fmt.Println("Pushed 1 commit to repository ", event.Repo.Name, "at", event.CreatedAt)
			} else {
				fmt.Println("Pushed", event.Payload.Size, "commits to repository ", event.Repo.Name, " at ", event.CreatedAt)
			}
		case "ReleaseEvent":
			fmt.Println("Published package on repository", event.Repo.Name, "at", event.CreatedAt)
		case "SponsorshipEvent":
			fmt.Println("Created sponsorship listing on repository", event.Repo.Name, "at", event.CreatedAt)
		case "WatchEvent":
			fmt.Println("Starred repository", event.Repo.Name, "at", event.CreatedAt)
		default:
			fmt.Println("Performed event", event.EventType, "on", event.Repo.Name, "at", event.CreatedAt)
		}
	}
}
