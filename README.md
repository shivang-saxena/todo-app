# Todo App

A simple Todo application written in Go.

## Features

- RESTful API for managing todo items
- CRUD operations (Create, Read, Update, Delete)
- PostgreSQL database for data persistence
- Clean architecture with separation of concerns

## Prerequisites

- Go 1.21 or higher
- PostgreSQL

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/shivang-saxena/todo-app.git
cd todo-app
```

### 2. Set up environment variables

Copy the example environment file and adjust as needed:

```bash
cp .env.example .env
```

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run main.go
```

The server will start at http://localhost:8080.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /api/v1/todos | Get all todos |
| GET    | /api/v1/todos/:id | Get a specific todo |
| POST   | /api/v1/todos | Create a new todo |
| PUT    | /api/v1/todos/:id | Update a todo |
| DELETE | /api/v1/todos/:id | Delete a todo |
| GET    | /health | Health check endpoint |

## Project Structure

```
.
├── config/          # Configuration (database, environment)
├── controllers/     # Request handlers
├── models/          # Database models
├── routes/          # API routes
├── .env.example     # Example environment variables
├── .gitignore       # Git ignore file
├── go.mod           # Go modules
├── main.go          # Application entry point
└── README.md        # This file
```

## Development

### Adding a New Feature

1. Create or modify the model(s) in the `models` directory
2. Add controller logic in the `controllers` directory
3. Register new routes in the `routes/router.go` file

## License

This project is open-source and available under the MIT License.
