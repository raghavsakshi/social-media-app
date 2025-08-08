package routes

import(
	"github.com/gofiber/fiber/v2"

	"social-media-app/controllers/posts"
)
func Posts(r fiber.Router){

	postsRoutes := r.Group("/users/:id/posts")
	postsRoutes.Post("/", posts.Add)
	
	postsRoutes.Get("/", posts.Get)
	postsRoutes.Delete("/:post_id", posts.Delete)
}