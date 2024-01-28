package routes

import (
	"github.com/devharshgupta/kafka-message-scheduler/service/scheduler"
	"github.com/gofiber/fiber/v2"
)

func SetSchedulerRoutes(route fiber.Router) {
	route.Post("/message", scheduler.ScheduleMessage)
}
