package main

import (
	"fizzbuzz_go/api/databases"
	"fizzbuzz_go/api/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// instanciate server
func StartApi() *fiber.App {
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

	return app
}

func init() {
	// init redis client
	databases.StartRedis()
}

func main() {
	app := StartApi()

	log.Fatal(app.Listen(":3000"))
}
