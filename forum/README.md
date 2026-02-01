# Forum Application


## Features

- **User Authentication**: Secure registration and login with bcrypt password hashing
-  **Post Management**: Create posts with titles, content, and multiple categories
-  **Comments**: Add comments to posts
-  **Voting System**: Like and dislike posts and comments
- **Categories**: Organise posts by categories
- **Filtering**: Filter posts by categories, user's own posts, or liked posts
- **Session Management**: Cookie-based sessions with UUID tokens
- **Modern UI**: Clean, dark-themed responsive design

## Technologies

- **Backend**: Go 1.24
- **Database**: SQLite3
- **Authentication**: bcrypt for password hashing, UUID for session tokens
- **Frontend**: HTML templates with modern CSS

## Installation

### Prerequisites

- Go 1.24 or higher
- Docker (optional)


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

