func LikedPostsHandler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user").(models.User)
    posts, err := models.GetLikedPosts(user.ID)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    renderTemplate(w, "index.html", map[string]interface{}{"Posts": posts})
}
