package model

import (
	"time"
)

type Subscription struct {
	Email string
	Name string
	CreatedAt time.Time
	Status int
}