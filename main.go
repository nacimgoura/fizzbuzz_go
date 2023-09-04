package main

import (
	"fizzbuzz_go/api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartApi() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "fizzbuzz",
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}?${queryParams}\n",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Routes
	router.SetupRoutes(app)

	// Start server
	app.Listen(":3010")
}

func main() {
	StartApi()
}
