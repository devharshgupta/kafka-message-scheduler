package models

import "time"

type Message struct {
	Id          string
	Key         string
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ScheduledAt time.Time
	IsPublished bool
}
