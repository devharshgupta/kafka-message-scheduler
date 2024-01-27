package router

import (
	"github.com/devharshgupta/kafka-message-scheduler/router/routes"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	root := app.Group("/")
	routes.SetRootRoutes(root)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // 404 route not found
	})

}
