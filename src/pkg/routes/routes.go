package routes

import (
	"my-fiber-app/mod/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    app.Get("/posts", controllers.GetPosts)
    app.Post("/posts", controllers.CreatePost)
}