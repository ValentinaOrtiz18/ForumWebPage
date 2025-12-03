package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	// POST request
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	err := database.CreateUser(email, username, string(hashed))
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Email already exists"})
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
