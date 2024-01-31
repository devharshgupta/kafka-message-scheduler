package structs

type SchedulerMessageBody struct {
	Id          string
	Priority    int
	Key         string                 `validate:"required"`
	Value       map[string]interface{} `validate:"required" gorm:"type:jsonb"`
	ScheduledAt string                 `validate:"required,timestamp"` //  RFC3339 time
}
