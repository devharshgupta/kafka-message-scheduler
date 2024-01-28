package structs

type SchedulerMessageBody struct {
	Id          string
	Priority    int
	Key         string `validate:"required"`
	Value       string `validate:"required"`
	ScheduledAt string `validate:"required"`
}
