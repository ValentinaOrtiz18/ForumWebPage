package handlers

import (
	"database/sql"
	"fmt"
	"forum/internal/database"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func getAuthenticatedUser(r *http.Request) (*database.User, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil, fmt.Errorf("not logged in")
	}

	userID, valid := database.GetUserIDBySession(cookie.Value)
	if !valid {
		return nil, fmt.Errorf("invalid session")
	}

	user, err := database.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return // <-- THIS FIXES THE PROBLEM
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// POST request continues here
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := database.GetUserByEmail(email)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	sessionID := uuid.New().String()
	expiration := time.Now().Add(2 * time.Hour)
	err = database.CreateSession(user.ID, sessionID, expiration)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  expiration,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_token")
	if cookie != nil {
		database.DeleteSession(cookie.Value)
		cookie.Expires = time.Now().Add(-time.Hour)
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
