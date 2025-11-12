func LoginHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        renderTemplate(w, "login.html", nil)
    case "POST":
        email := r.FormValue("email")
        password := r.FormValue("password")

        user, err := models.GetUserByEmail(email)
        if err != nil {
            renderTemplate(w, "login.html", map[string]string{"Error": "Invalid email or password"})
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
        if err != nil {
            renderTemplate(w, "login.html", map[string]string{"Error": "Invalid email or password"})
            return
        }

        sessionToken := uuid.NewString()
        expiration := time.Now().Add(24 * time.Hour)
        err = models.CreateSession(user.ID, sessionToken, expiration)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        cookie := &http.Cookie{
            Name:     "session_token",
            Value:    sessionToken,
            Expires:  expiration,
            HttpOnly: true,
        }
        http.SetCookie(w, cookie)
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_token")
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        user, err := models.GetUserBySession(cookie.Value)
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Add user info to context
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if err == nil {
        models.DeleteSession(cookie.Value)
        cookie.Expires = time.Now().Add(-time.Hour)
        http.SetCookie(w, cookie)
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}


