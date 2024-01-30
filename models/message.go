package models

type Message struct {
	Id          string
	Key         string
	Value       string
	ScheduledAt string
	IsPublished bool
}
