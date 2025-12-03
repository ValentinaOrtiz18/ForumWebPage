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
	http.HandleFunc("/", handlers.ShowLoginPage)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/post/create", handlers.PostHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}
