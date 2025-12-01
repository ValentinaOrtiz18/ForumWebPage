package main

import (
	"forum/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)
	http.HandleFunc("/posts", handlers.FilterPostsHandler)
	http.ListenAndServe(":8080", nil)
}
