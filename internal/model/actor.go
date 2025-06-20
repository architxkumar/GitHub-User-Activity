package model

type Actor struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	GravatarId string `json:"gravatar_id"`
	AvatarUrl  string `json:"avatar_url"`
	Url        string `json:"url"`
}
