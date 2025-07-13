package main

import (
	"embed"
	"my-fiber-app/mod/pkg/configs"
	"my-fiber-app/mod/pkg/routes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
    // Load environment variables
    configs.Init()

    // Create a new engine
	// engine := html.New("./views", ".html")

    // Or from an embedded system
    // See github.com/gofiber/embed for examples
   engine := html.NewFileSystem(http.Dir("./views"), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

    // engine := html.NewFileSystem(http.FS(viewsfs), ".html")
    // engine.AddFunc(
    //     // add unescape function
    //     "unescape", func(s string) template.HTML {
    //         return template.HTML(s)
    //     },
    // )
    // Pass the engine to the Views
    // app := fiber.New(fiber.Config{Views: engine})

    // Use logger middleware
    app.Use(logger.New())

    routes.Setup(app)

    app.Listen(":"+configs.AppPort()) // Start the server on the configured port
    // Additional setup can be done here
    // For example, setting up middleware, error handling, etc.
    // This is where you can add more routes or middleware as needed
    // You can also set up logging, CORS, or any other configurations required for your application
    // Ensure that the application is ready to handle requests
    // You can also add graceful shutdown logic here if needed
    // This is the entry point of your application, where everything comes together
    // You can also add health checks or other endpoints as needed
    // Make sure to handle any errors that might occur during initialization
}