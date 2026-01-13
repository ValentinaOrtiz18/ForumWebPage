package handlers

import (
	"html/template"
	"net/http"
	"forum/internal/database"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts []database.Post
	}{
		Posts: posts,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
