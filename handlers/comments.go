func CommentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    user := r.Context().Value("user").(models.User)
    postID, _ := strconv.Atoi(chi.URLParam(r, "postID"))
    content := r.FormValue("content")

    err := models.CreateComment(user.ID, postID, content)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
}
