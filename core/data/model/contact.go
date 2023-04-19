package model

import (
	"time"
)

type Contact struct {
	Id string
	Email string
	Phone string
	Name string
	Subject string
	Msg string
	CreatedAt time.Time
	Status int
}