package controllers

import (
    "github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
    // Logic to get posts
    return c.JSON([]Post{})
}

func CreatePost(c *fiber.Ctx) error {
    // Logic to create a post
    return c.JSON(Post{})
}