package main

import (
	"github.com/gofiber/fiber/v2"
	// "log"
	// "social-media-app/internals/cache"
	// "social-media-app/internals/database"
	// "social-media-app/internals/server"
	// "social-media-app/internals/notifications"
)


func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// This starts the server
	app.Listen(":3015")
}