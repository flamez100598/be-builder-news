package model

import "time"

type PageContent struct {
	Id int					`json:"id,omitempty"`
	Title string					`json:"title,omitempty"`
	Content string		`json:"content,omitempty"`
	UpdatedAt time.Time		`json:"updated_at,omitempty"`
	AvatarUpdatedBy string `json:"avatar_updated_by,omitempty"`
	UsernameUpdatedBy string `json:"username_updated_by,omitempty"`
	NameUpdatedBy string `json:"name_updated_by,omitempty"`
	EmailUpdatedBy string `json:"email_updated_by,omitempty"`
}

