
# Book Service

A REST ful service built with Go and the Gin framework for managing a library of books, integrated with Elasticsearch and Redis for robust searching and caching.

---

## Features

- **Book Management**: Add, update, delete, and retrieve books.
- **Advanced Search**: Search books by title, author, price range, and username.
- **User Activity Logging**: Log and display the last 3 user actions.
- **Caching**: Use Redis for caching user activity to enhance performance.
- **Elasticsearch Integration**: Efficiently manage and search data.

---

## Project Structure
```
book_service/
├── main.go                  # Entry point of the application
├── go.mod                   # Dependency management
├── go.sum                   # Dependency checksums
├── README.md                # Project documentation
├── config/
│   ├── consts.go            # Application constants
│   ├── structs.go           # Struct definitions
├── connectors/
│   ├── elastic_client.go    # Elasticsearch client setup
│   ├── redis_client.go      # Redis client setup
├── datastore/
│   ├── interfaces.go        # Interface definitions for data access
│   ├── elastic.go           # Elasticsearch implementation
│   ├── redis.go             # Redis implementation
├── middleware/
│   ├── log_user_action.go   # Middleware for logging user actions
├── model/
│   ├── book.go              # Book model
├── service/
│   ├── handler.go           # HTTP handlers
│   ├── routes.go            # API route definitions
├── utilities/
│   ├── utils.go             # Utility functions
```
---

## API Endpoints

### Books
- **POST /books**: Add a new book.
- **PUT /books**: Update book details.
- **GET /books**: Retrieve a book by ID.
- **DELETE /books**: Remove a book by ID.

### Search
- **GET /search**: Search books by title, author, price range, or username.

### Store Statistics
- **GET /store**: Retrieve the total number of books and distinct authors.

### Activity
- **GET /activity**: Retrieve the last 3 actions by a user.

---

## Dependencies

- **[Gin](https://github.com/gin-gonic/gin)**: A fast and lightweight web framework for Go.
- **[Elasticsearch](https://github.com/olivere/elastic)**: A distributed full-text search and analytics engine.
- **[Redis](https://github.com/go-redis/redis)**: An in-memory database for caching and data storage.

---
