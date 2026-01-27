package handlers

import (
	"fmt"
	"forum/internal/database"
	"net/http"
	"strconv"
)

// LikePostHandler handles liking a post
func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := getAuthenticatedUser(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = database.LikePost(user.ID, postID)
	if err != nil {
		http.Error(w, "Failed to like post", http.StatusInternalServerError)
		return
	}

	// Redirect back to the post
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = fmt.Sprintf("/post?id=%d", postID)
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

// DislikePostHandler handles disliking a post
func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := getAuthenticatedUser(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = database.DislikePost(user.ID, postID)
	if err != nil {
		http.Error(w, "Failed to dislike post", http.StatusInternalServerError)
		return
	}

	// Redirect back to the post
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = fmt.Sprintf("/post?id=%d", postID)
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

// LikeCommentHandler handles liking a comment
func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := getAuthenticatedUser(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	commentIDStr := r.FormValue("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	postIDStr := r.FormValue("post_id")

	err = database.LikeComment(user.ID, commentID)
	if err != nil {
		http.Error(w, "Failed to like comment", http.StatusInternalServerError)
		return
	}

	// Redirect back to the post
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = fmt.Sprintf("/post?id=%s", postIDStr)
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

// DislikeCommentHandler handles disliking a comment
func DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := getAuthenticatedUser(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	commentIDStr := r.FormValue("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	postIDStr := r.FormValue("post_id")

	err = database.DislikeComment(user.ID, commentID)
	if err != nil {
		http.Error(w, "Failed to dislike comment", http.StatusInternalServerError)
		return
	}

	// Redirect back to the post
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = fmt.Sprintf("/post?id=%s", postIDStr)
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}
