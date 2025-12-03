package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := database.GetPostByID(postID)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	comments, err := database.GetCommentsByPostID(postID)
	if err != nil {
		http.Error(w, "Could not load comments", http.StatusInternalServerError)
		return
	}

	categories, err := database.GetCategoriesByPostID(postID)
	if err != nil {
		http.Error(w, "Could not load categories", http.StatusInternalServerError)
		return
	}

	data := struct {
		Post       database.Post
		Comments   []database.Comment
		Categories []database.Category
	}{
		Post:       post,
		Comments:   comments,
		Categories: categories,
	}

	tmpl := template.Must(template.ParseFiles("templates/post.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
