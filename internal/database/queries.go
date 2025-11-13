type User struct {
    ID       int
    Email    string
    Username string
    Password string
}

func GetUserByEmail(email string) (User, error) {
    var u User
    err := DB.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).
        Scan(&u.ID, &u.Email, &u.Username, &u.Password)
    return u, err
}
