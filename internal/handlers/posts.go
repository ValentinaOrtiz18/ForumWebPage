package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Get post ID from query
	postIDStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Fetch post, comments, categories from DB
	post, _ := database.GetPostByID(postID)
	comments, _ := database.GetCommentsByPostID(postID)
	categories, _ := database.GetCategoriesByPostID(postID)

	data := struct {
		Post       database.Post
		Comments   []database.Comment
		Categories []string
	}{post, comments, categories}

	tmpl := template.Must(template.ParseFiles("templates/post.html"))
	tmpl.Execute(w, data)
}
