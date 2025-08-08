package routes

import (
	"social-media-app/controllers/friendships"
	"github.com/gofiber/fiber/v2"
)


func Friendships(r fiber.Router) {
	PostsRoutes := r.Group("/friends/:id/friends")
	PostsRoutes.Post("/", friendship.Add)
    PostsRoutes.Get("/:id", friendship.Get)
	PostsRoutes.Delete("/:id", friendship.Delete)

}