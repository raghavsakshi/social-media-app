package routes

import (
	"github.com/gofiber/fiber/v2"
	"social-media-app/controllers/users"
)

func Users(r fiber.Router) {
	userRoutes := r.Group("/users/:id/users")
	userRoutes.Post("/", users.Add)
  userRoutes.Get("/:id", users.GetAll)
	userRoutes.Get("/:id", users.Get)
	userRoutes.Delete("/:id", users.Delete)
}