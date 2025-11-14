package main

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
    ID        int
    PostID    int
    UserID    int
    Username  string
    Content   string
    CreatedAt time.Time
}

func CreateComment(userID, postID int, content string) error {
    _, err := db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)", userID, postID, content)
    return err
}

func GetCommentsByPostID(postID int) ([]Comment, error) {
    rows, err := db.Query(`
        SELECT c.id, c.post_id, c.user_id, u.username, c.content, c.created_at
        FROM comments c
        JOIN users u ON c.user_id = u.id
        WHERE c.post_id = ?
        ORDER BY c.created_at ASC`, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []Comment
    for rows.Next() {
        var c Comment
        rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Username, &c.Content, &c.CreatedAt)
        comments = append(comments, c)
    }
    return comments, nil
}
