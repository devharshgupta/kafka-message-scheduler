package models

type Message struct {
	Id          string
	Key         string                 `validate:"required"`
	Value       map[string]interface{} `validate:"required" gorm:"type:jsonb"`
	ScheduledAt string                 `validate:"required,timestamp"` //  RFC3339 time
	IsPublished bool
}
