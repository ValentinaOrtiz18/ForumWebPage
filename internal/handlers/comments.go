package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/internal/database"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    session, err := r.Cookie("session_token")
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    userID, valid := database.GetUserIDBySession(session.Value)
    if !valid {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    postID, _ := strconv.Atoi(r.FormValue("post_id"))
    content := r.FormValue("content")

    err = database.CreateComment(userID, postID, content)
    if err != nil {
        http.Error(w, "Failed to add comment", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postID), http.StatusSeeOther)
}

