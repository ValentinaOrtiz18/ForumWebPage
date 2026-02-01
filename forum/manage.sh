#!/bin/bash

# Forum Management Script

case "$1" in
    start)
        echo "🚀 Starting forum server..."
        go run cmd/main.go
        ;;
    
    build)
        echo "🔨 Building forum application..."
        go build -o forum ./cmd/main.go
        echo "✅ Build complete! Run with ./forum"
        ;;
    
    docker-build)
        echo "🐳 Building Docker image..."
        docker build -t forum-app .
        echo "✅ Docker image built! Run with: docker run -p 3030:3030 forum-app"
        ;;
    
    docker-up)
        echo "🐳 Starting forum with Docker Compose..."
        docker-compose up -d
        echo "✅ Forum is running at http://localhost:3030"
        ;;
    
    docker-down)
        echo "🐳 Stopping Docker containers..."
        docker-compose down
        echo "✅ Containers stopped"
        ;;
    
    test)
        echo "🧪 Running tests..."
        bash test.sh
        ;;
    
    clean)
        echo "🧹 Cleaning up..."
        rm -f forum forum.db
        echo "✅ Cleaned binary and database"
        ;;
    
    reset-db)
        echo "⚠️  Resetting database..."
        rm -f forum.db
        echo "✅ Database reset. Will be recreated on next run."
        ;;
    
    *)
        echo "Forum Management Script"
        echo ""
        echo "Usage: ./manage.sh [command]"
        echo ""
        echo "Commands:"
        echo "  start         - Start the forum server"
        echo "  build         - Build the forum binary"
        echo "  docker-build  - Build Docker image"
        echo "  docker-up     - Start with Docker Compose"
        echo "  docker-down   - Stop Docker containers"
        echo "  test          - Run endpoint tests"
        echo "  clean         - Remove binary and database"
        echo "  reset-db      - Reset the database"
        echo ""
        echo "Examples:"
        echo "  ./manage.sh start"
        echo "  ./manage.sh docker-up"
        echo "  ./manage.sh test"
        ;;
esac
