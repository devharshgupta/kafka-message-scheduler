package scheduler

import (
	"encoding/json"

	"github.com/devharshgupta/kafka-message-scheduler/common/kafka"
	"github.com/devharshgupta/kafka-message-scheduler/common/structs"
	"github.com/devharshgupta/kafka-message-scheduler/common/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ScheduleMessage(c *fiber.Ctx) error {

	data := new(structs.SchedulerMessageBody)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failure",
			"error":  "invalid body",
		})
	}

	errorMessages := validator.ValidateStruct(data)

	if len(errorMessages) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failure",
			"errors": errorMessages,
		})
	}

	if data.Id == "" {
		data.Id = uuid.New().String()
	}

	value, _ := json.Marshal(*data)

	err := kafka.PushSeheduledMessage(kafka.PushSeheduledMessageOptions{
		Key:      data.Key,
		Value:    string(value),
		Priority: data.Priority,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failure",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}
