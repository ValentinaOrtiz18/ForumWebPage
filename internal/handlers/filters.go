package handlers

import (
	"forum/internal/database"
	"net/http"
	"text/template"
)

func FilterPostsHandler(w http.ResponseWriter, r *http.Request) {
    filter := r.URL.Query().Get("filter")

    if filter == "liked" {
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

        posts, err := database.GetLikedPosts(userID)
        if err != nil {
            http.Error(w, "Could not load liked posts", http.StatusInternalServerError)
            return
        }

        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, posts)
        return
    }

    // default: show all posts
    posts, _ := database.GetAllPosts()
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, posts)
}
