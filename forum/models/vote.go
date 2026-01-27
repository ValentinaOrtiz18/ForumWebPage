package database

import (
	"database/sql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "database.DB")
	if err != nil {
		panic(err)
	}
}

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
}

func AddVote(userID, postID, commentID, voteType int) error {
	_, err := DB.Exec(`
        INSERT OR REPLACE INTO votes (user_id, post_id, comment_id, type)
        VALUES (?, ?, ?, ?)`, userID, postID, commentID, voteType)
	return err
}

func CountVotes(postID, commentID, voteType int) (int, error) {
	var count int
	err := DB.QueryRow(`
        SELECT COUNT(*) FROM votes WHERE 
        post_id = ? AND comment_id = ? AND type = ?`, postID, commentID, voteType).Scan(&count)
	return count, err
}

func GetLikedPosts(userID int) ([]Post, error) {
	rows, err := DB.Query(`
        SELECT p.id, p.title, p.content, p.created_at
        FROM posts p
        JOIN votes v ON v.post_id = p.id
        WHERE v.user_id = ? AND v.type = 1
        ORDER BY p.created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)
		posts = append(posts, p)
	}
	return posts, nil
}
