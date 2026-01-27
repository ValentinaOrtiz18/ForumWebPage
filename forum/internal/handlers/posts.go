package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
	"strconv"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Show create post form
		categories, _ := database.GetCategories()
		user, err := getAuthenticatedUser(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		data := struct {
			Categories []database.Category
			User       *database.User
			LoggedIn   bool
		}{
			Categories: categories,
			User:       user,
			LoggedIn:   true,
		}

		tmpl := template.Must(template.ParseFiles("templates/create_post.html"))
		tmpl.Execute(w, data)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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
	categoryIDs := r.Form["categories"] // Get multiple category IDs

	postID, err := database.CreatePost(userID, title, content)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Assign categories to the post
	for _, catIDStr := range categoryIDs {
		catID, err := strconv.Atoi(catIDStr)
		if err == nil {
			database.AssignCategoryToPost(int(postID), catID)
		}
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

	// Get like/dislike counts
	likes, _ := database.CountLikes(postID)
	dislikes, _ := database.CountDislikes(postID)

	// Get author
	author, _ := database.GetUserByID(post.UserID)

	// Check if user is logged in
	user, _ := getAuthenticatedUser(r)

	data := struct {
		Post       database.Post
		Comments   []database.Comment
		Categories []database.Category
		Likes      int
		Dislikes   int
		Author     database.User
		User       *database.User
		LoggedIn   bool
	}{
		Post:       post,
		Comments:   comments,
		Categories: categories,
		Likes:      likes,
		Dislikes:   dislikes,
		Author:     author,
		User:       user,
		LoggedIn:   user != nil,
	}

	tmpl := template.Must(template.ParseFiles("templates/post.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
