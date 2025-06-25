package model

// The repository object where the event occurred.
type repository struct {
	// The unique identifier of the repository.
	Id int `json:"id"`
	// The name of the repository, which includes the owner and repository name.
	//
	// For example, octocat/hello-world is the name of the hello-world repository owned by the octocat personal account.
	Name string `json:"name"`
	// The REST API URL used to retrieve the repository object, which includes additional repository information.
	Url string `json:"url"`
}
