package router

import (
	"fizzbuzz_go/api/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// get fizzbuzz with parameters
	app.Get("/", handler.FizzBuzz)

	// statistics endpoint allowing users to know what the most frequent request has been
	app.Get("/stats", handler.Stats)

	// 404  Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})
}
