package router

import (
	"fizzbuzz_go/api/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.FizzBuzz)

	// app.Get("/stats", handler.GetStats)

	// 404  Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})
}
