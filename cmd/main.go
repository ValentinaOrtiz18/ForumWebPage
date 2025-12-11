package main

import (
	"forum/internal/database"
	"forum/internal/handlers"
	"net/http"
)

func main() {
	// Initialize the database
	database.InitDB("forum.DB")

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/post", handlers.ViewPostHandler)
	http.HandleFunc("/post/create", handlers.CreatePostHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}
