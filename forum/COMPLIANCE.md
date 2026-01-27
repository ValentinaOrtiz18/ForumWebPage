# Forum Project - Compliance Checklist

## ✅ Project Requirements Verification

### Objectives
- [x] **Communication between users**: Users can create posts and comments
- [x] **Associating categories to posts**: Posts can have multiple categories
- [x] **Liking and disliking**: Both posts and comments can be liked/disliked
- [x] **Filtering posts**: Filter by categories, user posts, and liked posts

### SQLite Database
- [x] **SQLite usage**: Using github.com/mattn/go-sqlite3
- [x] **Entity relationships**: Proper foreign keys and relationships
- [x] **SELECT queries**: Used throughout for reading data
- [x] **CREATE queries**: schema.sql creates all tables
- [x] **INSERT queries**: Used for creating users, posts, comments, etc.

### Authentication
- [x] **User registration**: Email, username, and password required
- [x] **Email validation**: Checks for duplicate emails
- [x] **Username validation**: Checks for duplicate usernames
- [x] **Password encryption**: Using bcrypt (BONUS)
- [x] **Login session**: Cookie-based sessions
- [x] **Session expiration**: 2-hour expiration
- [x] **UUID for sessions**: Using google/uuid (BONUS)
- [x] **Credential verification**: Validates email and password

### Communication
- [x] **Post creation**: Only registered users can create posts
- [x] **Comment creation**: Only registered users can comment
- [x] **Category association**: Posts can have one or more categories
- [x] **Public viewing**: All users can view posts and comments
- [x] **Restricted actions**: Non-registered users cannot post/comment

### Likes and Dislikes
- [x] **Post voting**: Registered users can like/dislike posts
- [x] **Comment voting**: Registered users can like/dislike comments
- [x] **Vote counts visible**: All users can see like/dislike counts
- [x] **Unique constraints**: Users cannot vote multiple times

### Filtering
- [x] **Filter by categories**: Implemented via /filter?category=ID
- [x] **Filter by created posts**: /filter?filter=myposts (requires login)
- [x] **Filter by liked posts**: /filter?filter=liked (requires login)
- [x] **Subforum concept**: Categories work as subforums

### Docker
- [x] **Dockerfile**: Multi-stage build with Go and Alpine
- [x] **docker-compose.yml**: Easy deployment configuration
- [x] **Proper dependencies**: CGO enabled for SQLite support

### Error Handling
- [x] **HTTP status codes**: Proper status codes (400, 401, 404, 405, 500)
- [x] **Website errors**: User-friendly error pages
- [x] **Technical errors**: All errors are caught and handled
- [x] **Database errors**: Proper error handling in queries

### Code Quality
- [x] **Good practices**: Organized code structure
- [x] **Package organization**: Separated handlers, database, models
- [x] **No frontend frameworks**: Pure HTML/CSS/Go templates

### Allowed Packages
- [x] **Standard Go packages**: Only standard library used
- [x] **github.com/mattn/go-sqlite3**: For SQLite database
- [x] **golang.org/x/crypto/bcrypt**: For password hashing
- [x] **github.com/google/uuid**: For session tokens

## 📁 Project Structure

```
forum/
├── cmd/
│   └── main.go                    # Entry point, route setup
├── internal/
│   ├── database/
│   │   ├── queries.go            # All database operations
│   │   └── schema.sql            # Database schema
│   └── handlers/
│       ├── auth.go               # Login/logout/session management
│       ├── comments.go           # Comment creation
│       ├── filters.go            # Post filtering
│       ├── home.go               # Home page handler
│       ├── posts.go              # Post viewing and creation
│       ├── register.go           # User registration
│       └── votes.go              # Like/dislike handlers
├── models/                        # Data models (comment.go, vote.go)
├── static/
│   └── css/
│       └── style.css             # Modern minimalistic styling
├── templates/                     # HTML templates
│   ├── index.html                # Home/post list page
│   ├── post.html                 # Single post view
│   ├── create_post.html          # Create new post
│   ├── login.html                # Login page
│   └── register.html             # Registration page
├── Dockerfile                     # Docker configuration
├── docker-compose.yml            # Docker Compose setup
├── go.mod                        # Go dependencies
├── go.sum                        # Dependency checksums
├── README.md                     # Project documentation
├── test.sh                       # Basic endpoint tests
└── .gitignore                    # Git ignore rules
```

## 🎨 Features Implemented

### Core Features
1. **User Management**
   - Registration with email, username, password
   - Secure login with bcrypt
   - Session management with UUID tokens
   - 2-hour session expiration

2. **Post System**
   - Create posts with title and content
   - Assign multiple categories
   - View all posts or filtered posts
   - Post metadata (author, timestamp)

3. **Comment System**
   - Add comments to posts
   - View all comments
   - Comment metadata

4. **Voting System**
   - Like/dislike posts
   - Like/dislike comments
   - Real-time vote counts
   - Prevents duplicate votes

5. **Category System**
   - 8 default categories (General, Technology, Programming, Gaming, Science, Entertainment, Sports, News)
   - Filter posts by category
   - Multiple categories per post

6. **Filtering**
   - View all posts
   - View only your posts
   - View liked posts
   - View posts by category

### Bonus Features
- ✅ Password encryption with bcrypt
- ✅ UUID session tokens
- ✅ Modern, responsive UI
- ✅ Comment voting system
- ✅ Default categories
- ✅ Docker support
- ✅ Docker Compose configuration

## 🚀 Quick Start

### Run Locally
```bash
go run cmd/main.go
```

### Run with Docker
```bash
docker-compose up -d
```

### Run Tests
```bash
bash test.sh
```

## 📊 Database Schema

### Tables
- `users` - User accounts (id, email, username, password)
- `sessions` - Active sessions (id, user_id, session_token, expires_at)
- `posts` - Forum posts (id, user_id, title, content, created_at)
- `comments` - Post comments (id, post_id, user_id, content, created_at)
- `categories` - Post categories (id, name)
- `post_categories` - Many-to-many post-category relationship
- `likes` - Post likes (id, user_id, post_id, created_at)
- `dislikes` - Post dislikes (id, user_id, post_id, created_at)
- `comment_likes` - Comment likes (id, user_id, comment_id, created_at)
- `comment_dislikes` - Comment dislikes (id, user_id, comment_id, created_at)

## 🎯 All Requirements Met

✅ This project meets 100% of the specified requirements:
- SQLite database with proper schema
- User authentication with cookies and sessions
- Post and comment creation (registered users only)
- Category system
- Like/dislike functionality
- Filtering by category, user posts, and liked posts
- Docker support
- Proper error handling
- Good code practices
- Only allowed packages used
- No frontend frameworks
