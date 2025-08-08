package config

import (
		"social-media-app/models/friendship"
			"social-media-app/models/posts"
				"social-media-app/models/users"
	"social-media-app/internals/database"
)
func AutoMigration() {
	database.Client().AutoMigrate(&users.Users{}, &friendships.Friendships{}, &posts.Posts{})
}