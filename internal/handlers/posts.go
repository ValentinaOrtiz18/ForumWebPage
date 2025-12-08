package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
	"strconv"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	userID, valid := database.GetUserIDBySession(cookie.Value)
	if !valid {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err = database.CreatePost(userID, title, content)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ViewPostHandler(w http.ResponseWriter, r *http.Request) {
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
