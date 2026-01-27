package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get posts
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	// Get categories
	categories, err := database.GetCategories()
	if err != nil {
		categories = []database.Category{}
	}

	// Check if user is logged in
	user, _ := getAuthenticatedUser(r)

	data := struct {
		Posts      []database.Post
		Categories []database.Category
		User       *database.User
		LoggedIn   bool
	}{
		Posts:      posts,
		Categories: categories,
		User:       user,
		LoggedIn:   user != nil,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
