package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
	"strconv"
)

// FilterPostsHandler handles displaying posts filtered by type:
// "all" = all posts, "myposts" = posts by the logged-in user, "liked" = posts liked by the user, or by category ID
func FilterPostsHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")          // "all", "myposts", or "liked"
	categoryIDStr := r.URL.Query().Get("category") // category ID
	var posts []database.Post
	var err error

	// Get userID from session if needed
	var userID int
	var loggedIn bool
	user, authErr := getAuthenticatedUser(r)
	if authErr == nil {
		userID = user.ID
		loggedIn = true
	}

	// Get all categories for sidebar
	categories, _ := database.GetCategories()

	switch {
	case categoryIDStr != "":
		// Filter by category
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		posts, err = database.GetPostsByCategory(categoryID)
		if err != nil {
			http.Error(w, "Failed to load posts by category", http.StatusInternalServerError)
			return
		}

	case filter == "myposts":
		if !loggedIn {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetUserPosts(userID)
		if err != nil {
			http.Error(w, "Failed to load your posts", http.StatusInternalServerError)
			return
		}

	case filter == "liked":
		if !loggedIn {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetLikedPosts(userID)
		if err != nil {
			http.Error(w, "Failed to load liked posts", http.StatusInternalServerError)
			return
		}

	default: // "all" or empty
		posts, err = database.GetAllPosts()
		if err != nil {
			http.Error(w, "Failed to load posts", http.StatusInternalServerError)
			return
		}
	}

	// Wrap posts in a struct for the template
	data := struct {
		Posts      []database.Post
		Filter     string
		Categories []database.Category
		User       *database.User
		LoggedIn   bool
	}{
		Posts:      posts,
		Filter:     filter,
		Categories: categories,
		User:       user,
		LoggedIn:   loggedIn,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
	}
}
