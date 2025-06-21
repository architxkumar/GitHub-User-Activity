package model

type payload struct {
	// The type of Git ref object created in the repository.
	// Can be either branch, tag, or repository.
	// Note: For usage with "CreateEvent" type
	ReferenceType string `json:"ref_type"`
	Size          int    `json:"size"`
	Action        string `json:"action"`
}
