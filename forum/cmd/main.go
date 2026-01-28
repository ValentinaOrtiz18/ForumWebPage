package main

import (
	"forum/internal/database"
	"forum/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialise the database
	database.InitDB("./forum.db")

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/post", handlers.ViewPostHandler)
	http.HandleFunc("/post/create", handlers.CreatePostHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)
	http.HandleFunc("/post/like", handlers.LikePostHandler)
	http.HandleFunc("/post/dislike", handlers.DislikePostHandler)
	http.HandleFunc("/comment/like", handlers.LikeCommentHandler)
	http.HandleFunc("/comment/dislike", handlers.DislikeCommentHandler)
	http.HandleFunc("/filter", handlers.FilterPostsHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
