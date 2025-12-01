package main

import (
	"forum/internal/database"
	"forum/internal/handlers"
	"net/http"
)

func main() {
	// Initialize the database
	database.InitDB("forum.db")

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/post/create", handlers.PostHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}
