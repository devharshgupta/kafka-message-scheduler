package routes

import "github.com/gofiber/fiber/v2"

func SetRootRoutes(root fiber.Router) {
	root.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
