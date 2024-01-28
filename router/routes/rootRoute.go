package routes

import "github.com/gofiber/fiber/v2"

func SetRootRoutes(route fiber.Router) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"data":   "Hello 👋 World 🗺️",
		})
	})
}
