// main.go
package main

import (
	"github.com/devharshgupta/kafka-message-scheduler/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router.InitRoutes(app)

	app.Listen(":3000")
}
