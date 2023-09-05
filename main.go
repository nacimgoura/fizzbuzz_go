package main

import (
	"fizzbuzz_go/api/databases"
	"fizzbuzz_go/api/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// instanciate server
func StartApi() *fiber.App {
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
	host := "localhost:6379"
	// override host if env var REDIS_HOST is set
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}
	// init redis client
	databases.StartRedis(host)
}

func main() {
	app := StartApi()

	log.Fatal(app.Listen(":3000"))
}
