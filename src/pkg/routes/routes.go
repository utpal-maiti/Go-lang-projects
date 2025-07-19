package routes

import (
	"my-fiber-app/mod/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all routes for the Fiber application.
func Setup(app *fiber.App) {
	// Initialize routes for the application
	// This function sets up the routes for the Fiber application
	// Define the base route
	app.Get("/", func(c *fiber.Ctx) error {
			// Render index
			return c.Render("index", fiber.Map{
				"Title": "Hello, World!",
				"content": "/",
			})
		})

	// Define a route for rendering the index template with a layout
	app.Get("/layout", func(c *fiber.Ctx) error {
			// Render index within layouts/main
			return c.Render("index", fiber.Map{
				"Title": "Hello, World!",
				"content": "/layout",
			}, "layouts/main")
		})

	app.Get("/layouts-nested", func(c *fiber.Ctx) error {
		// Render index within layouts/nested/main within layouts/nested/base
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"content": "/layouts-nested",
		}, "layouts/nested/main", "layouts/nested/base")
	})
	
	// This is the main entry point for the application
	// Define the base API route
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the API"})
	})
	// Define the versioned API route
	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the API version 1"})
	})
	// Define routes for post operations
	app.Get("/api/v1/posts", controllers.GetPosts)
	// Define routes for individual post operations
	app.Get("/api/v1/posts/:id", controllers.GetPostByID)
	// Define routes for creating, updating, and deleting posts 
	app.Post("/api/v1/posts", controllers.CreatePost)
	// Define route for updating a post
	app.Put("/api/v1/posts", controllers.UpdatePost)
	// Define routes for updating and deleting posts by ID
	app.Put("/api/v1/posts/:id", controllers.UpdatePost)
	// Define route for deleting a post by ID
	app.Delete("/api/v1/posts/:id", controllers.DeletePost)
	app.Get("/download", controllers.DownloadFile)

}
