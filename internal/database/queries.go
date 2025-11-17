package database

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type Post struct {
	ID        int
	UserID    int
	Title     string
	Content   string
	CreatedAt time.Time
}

func GetUserByEmail(email string) (User, error) {
	var u User
	err := DB.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).
		Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	return u, err
}

func CreateSession(UserID int, sessionToken string, expires time.Time) error {
	_, err := DB.Exec(`
		INSERT INTO sessions (user_id, session_token, expires_at)
		VALUES (?, ?, ?)
	`, UserID, sessionToken, expires)

	return err
}

func DeleteSession(token string) error {
	_, err := DB.Exec(`
		DELETE FROM sessions WHERE session_token = ?
	`, token)

	return err
}

func CreateComment(UserID, postID int, content string) error {
	_, err := DB.Exec(`
    INSERT INTO comments (user_id, post_id, content, created_at)
    VALUES (?, ?, ?, CURRENT_TIMESTAMP)
`, UserID, postID, content)

	return err
}

func GetUserIDBySession(token string) (int, bool) {
	var userID int
	err := DB.QueryRow(`
		SELECT user_id
		FROM sessions
		WHERE session_token = ? AND expires_at > CURRENT_TIMESTAMP
	`, token).Scan(&userID)

	if err != nil {
		return 0, false
	}

	return userID, true
}

func GetAllPosts() ([]Post, error) {
	rows, err := DB.Query(`
		SELECT id, user_id, title, content, created_at
		FROM posts
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func GetLikedPosts(userID int) ([]Post, error) {
	rows, err := DB.Query(`
		SELECT p.id, p.user_id, p.title, p.content, p.created_at
		FROM posts p
		JOIN likes l ON p.id = l.post_id
		WHERE l.user_id = ?
		ORDER BY p.created_at DESC
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
