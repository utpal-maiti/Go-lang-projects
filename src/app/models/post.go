package models

// Post represents a blog post with an ID, title, and body.
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}