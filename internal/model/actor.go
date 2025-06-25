package model

// Actor represents the user that triggered the event
//
// For more information, visit: (https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#event-object-common-properties)
type Actor struct {
	// Id is the unique identifier for the actor
	Id int `json:"id"`
	// Login is username of the actor
	Login string `json:"login"`
	// GravatarId is the unique identifier of the Gravatar profile for the actor
	GravatarId string `json:"gravatar_id"`
	// AvatarUrl is the URL of the actor's profile image.
	AvatarUrl string `json:"avatar_url"`
	// Url is the REST API URL used to retrieve the user object, which includes additional user information.
	Url string `json:"url"`
}
