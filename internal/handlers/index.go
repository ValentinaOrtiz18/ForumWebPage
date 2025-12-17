package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	var userID int
	var loggedIn bool

	cookie, err := r.Cookie("session_token")
	if err == nil {
		userID, loggedIn = database.GetUserIDBySession(cookie.Value)
	}

	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, "Could not load posts", http.StatusInternalServerError)
		return
	}

	data := struct {
		LoggedIn bool
		UserID   int
		Posts    []database.Post
	}{
		LoggedIn: loggedIn,
		UserID:   userID,
		Posts:    posts,
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
