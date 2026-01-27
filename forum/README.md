# Forum Application

A modern, minimalistic web forum built with Go and SQLite.

## Features

- 🔐 **User Authentication**: Secure registration and login with bcrypt password hashing
- 📝 **Post Management**: Create posts with titles, content, and multiple categories
- 💬 **Comments**: Add comments to posts
- 👍👎 **Voting System**: Like and dislike posts and comments
- 🏷️ **Categories**: Organize posts by categories
- 🔍 **Filtering**: Filter posts by categories, user's own posts, or liked posts
- 🍪 **Session Management**: Cookie-based sessions with UUID tokens
- 🎨 **Modern UI**: Clean, dark-themed responsive design

## Technologies

- **Backend**: Go 1.24
- **Database**: SQLite3
- **Authentication**: bcrypt for password hashing, UUID for session tokens
- **Frontend**: HTML templates with modern CSS

## Installation

### Prerequisites

- Go 1.24 or higher
- Docker (optional, for containerized deployment)

### Local Setup

1. Clone the repository:
```bash
cd forum
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

4. Access the forum at: `http://localhost:3030`

### Docker Setup

1. Build and run with Docker Compose:
```bash
docker-compose up -d
```

Or manually:

1. Build the Docker image:
```bash
docker build -t forum-app .
```

2. Run the container:
```bash
docker run -p 3030:3030 forum-app
```

3. Access the forum at: `http://localhost:3030`

To stop the container (if using Docker Compose):
```bash
docker-compose down
```

## Project Structure

```
forum/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── database/
│   │   ├── queries.go       # Database operations
│   │   └── schema.sql       # Database schema
│   └── handlers/
│       ├── auth.go          # Authentication handlers
│       ├── comments.go      # Comment handlers
│       ├── filters.go       # Filter handlers
│       ├── home.go          # Home page handler
│       ├── posts.go         # Post handlers
│       ├── register.go      # Registration handler
│       └── votes.go         # Like/dislike handlers
├── models/                  # Data models
├── static/
│   └── css/
│       └── style.css        # Application styles
├── templates/               # HTML templates
│   ├── index.html          # Home page
│   ├── post.html           # Post detail page
│   ├── create_post.html    # Create post form
│   ├── login.html          # Login page
│   └── register.html       # Registration page
├── go.mod                   # Go dependencies
├── go.sum                   # Go checksums
└── Dockerfile              # Docker configuration
```

## Features Details

### Authentication
- Secure user registration with email validation
- Password encryption using bcrypt
- Session-based authentication with UUID tokens
- 2-hour session expiration

### Posts
- Create posts with title and content
- Assign multiple categories to posts
- View all posts or filter by specific criteria
- See post author and creation date

### Comments
- Add comments to any post
- View all comments on a post
- Comments show author and timestamp

### Voting
- Like or dislike posts
- Like or dislike comments
- View vote counts (visible to all users)
- Voting restricted to registered users

### Filtering
- View all posts
- Filter by category
- View only your own posts (requires login)
- View posts you've liked (requires login)

## Database Schema

The application uses SQLite with the following tables:
- `users` - User accounts
- `sessions` - Active user sessions
- `posts` - Forum posts
- `comments` - Post comments
- `categories` - Post categories
- `post_categories` - Post-category relationships
- `likes` - Post likes
- `dislikes` - Post dislikes
- `comment_likes` - Comment likes
- `comment_dislikes` - Comment dislikes

## API Endpoints

- `GET /` - Home page (list all posts)
- `GET /login` - Login page
- `POST /login` - Login handler
- `GET /register` - Registration page
- `POST /register` - Registration handler
- `GET /logout` - Logout handler
- `GET /post` - View single post
- `GET /post/create` - Create post form
- `POST /post/create` - Create post handler
- `POST /post/like` - Like post
- `POST /post/dislike` - Dislike post
- `POST /comment/create` - Create comment
- `POST /comment/like` - Like comment
- `POST /comment/dislike` - Dislike comment
- `GET /filter` - Filter posts

## License

This project is created for educational purposes.
