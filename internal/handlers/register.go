package handlers

import (
	"forum/internal/database"
	"html/template"
	"net/http"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	// ✅ ALLOW GET (this fixes the 405)
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	// ❌ Block anything that is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// POST logic
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	err = database.CreateUser(email, username, password)
	if err != nil {
		tmpl.Execute(w, map[string]string{
			"Error": "Registration failed",
		})
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
