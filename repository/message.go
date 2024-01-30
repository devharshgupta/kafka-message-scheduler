package repository

import (
	"github.com/devharshgupta/kafka-message-scheduler/common/db"
	"github.com/devharshgupta/kafka-message-scheduler/models"
)

func CreateMessages(data []models.Message) error {
	db := db.DB

	result := db.Create(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
