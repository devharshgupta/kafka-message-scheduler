// main.go
package main

import (
	"log"
	"os"

	"github.com/devharshgupta/kafka-message-scheduler/common/constant"
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
	kafka.EnsureKafkaTopicsExitInit(constant.KAFKA_TOPICS)
	go kafka.InitConsumer(constant.KAFKA_TOPICS, 100) // polling @100ms
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router.InitRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
