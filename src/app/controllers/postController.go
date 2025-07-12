package controllers

import (
	"my-fiber-app/mod/app/models"
	"strconv"

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

// Handler to get a single post by ID
func GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	// TODO: Add logic to fetch post by id from database or data source
	// Example response:
	return c.JSON(fiber.Map{
		"id":      id,
		"message": "Single post fetched successfully",
	})
}

func CreatePost(c *fiber.Ctx) error {
    // Logic to create a post
    var post models.Post
    if err := c.BodyParser(&post); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(post)
}
func UpdatePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post
    if err := c.BodyParser(&post); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    // Convert id from string to int
    intID, err := strconv.Atoi(id)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }
    post.ID = intID // Assign converted int ID
    return c.JSON(post)
}   
func DeletePost(c *fiber.Ctx) error {
    id := c.Params("id")
    // Logic to delete a post
    // Convert id from string to int
    intID, err := strconv.Atoi(id)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }
    return c.JSON(fiber.Map{"message": "Post deleted", "id": intID})
}
