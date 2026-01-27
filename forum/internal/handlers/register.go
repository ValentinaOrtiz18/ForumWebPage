package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// ✅ DO NOT HASH HERE
	err := database.CreateUser(email, username, password)
	if err != nil {
		tmpl.Execute(w, map[string]string{
			"Error": "Registration failed",
		})
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
