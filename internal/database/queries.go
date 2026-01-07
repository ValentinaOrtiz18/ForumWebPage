package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

// ------------------- Structs -------------------

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

type Comment struct {
	ID        int
	PostID    int
	UserID    int
	Content   string
	CreatedAt time.Time
}

type Category struct {
	ID   int
	Name string
}

// ------------------- Initialization -------------------

func InitDB(path string) *sql.DB {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return DB
}

func CreateUser(email, username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        INSERT INTO users (email, username, password) VALUES (?, ?, ?)
    `, email, username, string(hashed))

	return err
}

func GetUserByEmail(email string) (User, error) {
	var u User
	err := DB.QueryRow(`
        SELECT id, email, username, password FROM users WHERE email = ?
    `, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	return u, err
}

func GetUserByID(id int) (User, error) {
	var u User
	err := DB.QueryRow(`
        SELECT id, email, username, password FROM users WHERE id = ?
    `, id).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	return u, err
}

// ------------------- Session Functions -------------------

func CreateSession(userID int, token string, expires time.Time) error {
	_, err := DB.Exec(`
        INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)
    `, userID, token, expires)
	return err
}

func DeleteSession(token string) error {
	_, err := DB.Exec(`DELETE FROM sessions WHERE session_token = ?`, token)
	return err
}

func GetUserIDBySession(token string) (int, bool) {
	var userID int
	err := DB.QueryRow(`
        SELECT user_id FROM sessions WHERE session_token = ? AND expires_at > CURRENT_TIMESTAMP
    `, token).Scan(&userID)
	if err != nil {
		return 0, false
	}
	return userID, true
}

// ------------------- Post Functions -------------------

func CreatePost(userID int, title, content string) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)
    `, userID, title, content)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetPostByID(id int) (Post, error) {
	var p Post
	err := DB.QueryRow(`
        SELECT id, user_id, title, content, created_at FROM posts WHERE id = ?
    `, id).Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt)
	return p, err
}

func GetAllPosts() ([]Post, error) {
	rows, err := DB.Query(`SELECT id, user_id, title, content, created_at FROM posts ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// ------------------- Comment Functions -------------------

func CreateComment(userID, postID int, content string) error {
	_, err := DB.Exec(`
        INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)
    `, userID, postID, content)
	return err
}

func GetCommentsByPostID(postID int) ([]Comment, error) {
	rows, err := DB.Query(`
        SELECT id, post_id, user_id, content, created_at
        FROM comments
        WHERE post_id = ?
        ORDER BY created_at ASC
    `, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

// ------------------- Likes / Dislikes Functions -------------------
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

func DislikePost(userID, postID int) error {
	_, err := DB.Exec(`INSERT OR IGNORE INTO dislikes (user_id, post_id) VALUES (?, ?)`, userID, postID)
	// remove like if exists
	DB.Exec(`DELETE FROM likes WHERE user_id = ? AND post_id = ?`, userID, postID)
	return err
}

func CountLikes(postID int) (int, error) {
	var count int
	err := DB.QueryRow(`SELECT COUNT(*) FROM likes WHERE post_id = ?`, postID).Scan(&count)
	return count, err
}

func CountDislikes(postID int) (int, error) {
	var count int
	err := DB.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE post_id = ?`, postID).Scan(&count)
	return count, err
}

// ------------------- Category Functions -------------------

func CreateCategory(name string) (int64, error) {
	res, err := DB.Exec(`INSERT OR IGNORE INTO categories (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetCategories() ([]Category, error) {
	rows, err := DB.Query(`SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cats := []Category{}
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

func AssignCategoryToPost(postID, categoryID int) error {
	_, err := DB.Exec(`INSERT OR IGNORE INTO post_categories (post_id, category_id) VALUES (?, ?)`, postID, categoryID)
	return err
}

func GetPostsByCategory(categoryID int) ([]Post, error) {
	rows, err := DB.Query(`
        SELECT p.id, p.user_id, p.title, p.content, p.created_at
        FROM posts p
        JOIN post_categories pc ON p.id = pc.post_id
        WHERE pc.category_id = ?
        ORDER BY p.created_at DESC
    `, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func GetCategoriesByPostID(postID int) ([]Category, error) {
	rows, err := DB.Query(`
		SELECT c.id, c.name
		FROM categories c
		JOIN post_categories pc ON c.id = pc.category_id
		WHERE pc.post_id = ?
	`, postID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var cat Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func GetUserPosts(userID int) ([]Post, error) {
	rows, err := DB.Query(`
		SELECT id, user_id, title, content, created_at
		FROM posts
		WHERE user_id = ?
		ORDER BY created_at DESC
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
