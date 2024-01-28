package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func RequestBodyValidator(ctx *fiber.Ctx) error {
	// Check if the request has a JSON Content-Type
	contentType := ctx.Get("Content-Type")
	if contentType != "application/json" {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"status": "failure",
			"error":  "Unsupported Media Type. Expecting application/json",
		})
	}

	isValid := json.Valid(ctx.Request().Body())

	if !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failure",
			"error":  "invalid request",
		})
	}

	return ctx.Next()
}
