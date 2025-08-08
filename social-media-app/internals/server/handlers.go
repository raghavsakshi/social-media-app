package server

import (
	"social-media-app/routes"

	"github.com/gofiber/fiber/v2"
)

//default error handler
func ErrorHandler (c *fiber.Ctx, e error) error {
	msg := e.Error()
	c.Status(fiber.StatusInternalServerError).JSON(msg)
	return nil
}
var notFoundHandler = func (c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Requested resource not found",
	})
}
func addRoutes(app *fiber.App) {
	baseRouter := app.Group("/social-media-app")
	routes.Users(baseRouter)
	routes.Friendships(baseRouter)
	routes.Posts(baseRouter)
}
