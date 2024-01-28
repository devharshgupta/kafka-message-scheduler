package router

import (
	"github.com/devharshgupta/kafka-message-scheduler/common/middleware"
	"github.com/devharshgupta/kafka-message-scheduler/router/routes"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	// validate request for valid json body
	app.Use("POST", "PATCH", "PUT", "DELETE", middleware.RequestBodyValidator)

	root := app.Group("/")
	routes.SetRootRoutes(root)

	scheduler := app.Group("/v1/scheduler")
	routes.SetSchedulerRoutes(scheduler)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // 404 route not found
	})

}
