package model

type User struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}