package routes

import "github.com/gofiber/fiber/v2"
import "github.com/samadguliev/posts-app/backend/handlers"

func SetupRoutes(app *fiber.App) {
	app.Get("posts", handlers.PostList)
	app.Post("posts", handlers.AddPost)
}
