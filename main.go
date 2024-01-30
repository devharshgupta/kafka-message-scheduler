// main.go
package main

import (
	"log"
	"os"

	"github.com/devharshgupta/kafka-message-scheduler/common/db"
	"github.com/devharshgupta/kafka-message-scheduler/common/env"
	"github.com/devharshgupta/kafka-message-scheduler/common/kafka"
	"github.com/devharshgupta/kafka-message-scheduler/common/validator"
	"github.com/devharshgupta/kafka-message-scheduler/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	env.InitEnv()
	db.InitDatabase()
	validator.InitValidator()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router.InitRoutes(app)

	go kafka.InitConsumer([]string{"Kafka_message_scheduler.P-1_Message", "Kafka_message_scheduler.P0_Message", "Kafka_message_scheduler.P1_Message"}, 100) // polling @100ms

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))

}
