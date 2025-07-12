package controllers

import (
	"my-fiber-app/mod/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
    // Logic to get posts
    posts := []models.Post{
        {ID: 1, Title: "First Post", Body: "This is the content of the first post."},
        {ID: 2, Title: "Second Post", Body: "This is the content of the second post."},
    }
    return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
    // Logic to create a post
    var post models.Post
    if err := c.BodyParser(&post); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(post)
}