# Basic Web API

A simple RESTful API built with Go and PostgreSQL that manages users and posts.

## Project Structure

```
basic_web_api/
├── db/
│   ├── db.go         # Database connection and initialization
│   └── db_test.go    # Database tests
├── handlers/
│   └── handlers.go   # HTTP request handlers
├── models/
│   └── user.go       # Data models (User and Post)
├── repository/
│   └── repository.go # Database operations
├── router/
│   └── router.go     # HTTP routing
├── utils/
│   └── date.go       # Utility functions
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
└── main.go         # Application entry point
```

## Prerequisites

- Go 1.20 or later
- PostgreSQL 16
- Docker (for running PostgreSQL)

## Setup

1. Start the PostgreSQL container:
```bash
docker run --name basic-web-api-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=basic_web_api -p 5432:5432 -d postgres:16
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on port 8080, and the database tables will be automatically created.

## API Endpoints

### Users

#### GET /users
Retrieves all users.

Response:
```json
[
    {
        "id": 1,
        "name": "John Doe",
        "dob": "1990-01-01T00:00:00Z"
    }
]
```

#### POST /users
Creates a new user.

Request:
```json
{
    "name": "John Doe",
    "dob": "1990-01-01"
}
```

Response:
```json
{
    "id": 1,
    "name": "John Doe",
    "dob": "1990-01-01T00:00:00Z"
}
```

### Posts

#### GET /posts
Retrieves all posts.

Response:
```json
[
    {
        "id": 1,
        "content": "Hello, World!",
        "created_at": "2025-05-14T10:00:00Z"
    }
]
```

#### POST /posts
Creates a new post.

Request:
```json
{
    "content": "Hello, World!"
}
```

Response:
```json
{
    "id": 1,
    "content": "Hello, World!",
    "created_at": "2025-05-14T10:00:00Z"
}
```

## Project Components

### Database (db/)
- Manages PostgreSQL connection
- Handles database initialization
- Creates required tables on startup
- Implements retry logic for connection stability

### Models (models/)
- Defines data structures for Users and Posts
- Implements custom JSON marshaling for date handling
- Includes database tags for mapping

### Repository (repository/)
- Implements data access layer
- Handles database operations for Users and Posts
- Provides clean separation from business logic

### Handlers (handlers/)
- HTTP request handlers for each endpoint
- Request validation and response formatting
- Error handling

### Router (router/)
- URL routing and request mapping
- Simple request routing based on path

### Utils (utils/)
- Utility functions for date parsing
- Common helper functions

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Successful GET requests
- 201: Successful POST requests
- 400: Invalid request data
- 404: Resource not found
- 500: Server error

## Future Improvements

Potential enhancements for the project:
1. Authentication and authorization
2. Input validation
3. Request logging
4. Environment configuration
5. API versioning
6. API Documentation
The API is documented using Swagger/OpenAPI 3 specification. The documentation includes:

- Detailed endpoint descriptions
- Request/response schemas
- Example requests and responses
- Error responses and status codes
- Data models and their relationships

To view the API documentation:

1. Generate the Swagger docs:
   ```bash
   swag init
   ```

2. Documentation files:
   - `docs/swagger.yaml`: OpenAPI 3.0 specification
   - `docs/swagger.json`: JSON version of the API specification
   - `docs/docs.go`: Go code generated from the spec

The documentation covers:
- User Management API
  - GET /users: List all users
  - POST /users: Create a new user
- Post Management API
  - GET /posts: List all posts
  - POST /posts: Create a new post

Each endpoint is documented with:
- Description
- Request parameters
- Response format
- Example payloads
- Possible error responses

Future Documentation Improvements:
- Add interactive Swagger UI
- Include authentication documentation
- Add rate limiting information
- Implement versioning
- Add automated documentation updates via CI/CD
7. More comprehensive error handling
8. Unit and integration tests

## Authorization

The API uses JWT (JSON Web Token) for authentication. To access protected endpoints (/users and /posts), you must first register or login to obtain a token.

### Register a New User

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "secretpass123",
    "dob": "1990-01-01"
  }'
```

Response:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
}
```

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "secretpass123"
  }'
```

Response:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
}
```

### Using the Token

Use the token in the Authorization header for protected endpoints:

```bash
# Get all users
curl -X GET http://localhost:8080/users \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."

# Create a new post
curl -X POST http://localhost:8080/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..." \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Hello, World!"
  }'
```

Note: Replace `eyJhbGciOiJIUzI1NiIs...` with the actual token received from register/login.
