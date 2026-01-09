package handlers

import (
	"net/http"
	"text/template"

	"forum/internal/database"
)

// FilterPostsHandler handles displaying posts filtered by type:
// "all" = all posts, "myposts" = posts by the logged-in user, "liked" = posts liked by the user
func FilterPostsHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter") // "all", "myposts", or "liked"
	var posts []database.Post
	var err error

	// Get userID from session if needed
	var userID int
	var loggedIn bool
	cookie, cookieErr := r.Cookie("session_token")
	if cookieErr == nil {
		userID, loggedIn = database.GetUserIDBySession(cookie.Value)
	}

	switch filter {

	case "myposts":
		if !loggedIn {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetUserPosts(userID) // You need this function in your DB
		if err != nil {
			http.Error(w, "Failed to load your posts", http.StatusInternalServerError)
			return
		}

	case "liked":
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
		Posts    []database.Post
		Filter   string
		LoggedIn bool
	}{
		Posts:    posts,
		Filter:   filter,
		LoggedIn: loggedIn,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
	}
}
