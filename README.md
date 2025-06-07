# Go TODO App

This is a simple TODO web application built for learning purposes. The project follows a clean architecture approach using the [Echo](https://echo.labstack.com/) framework and MySQL as the database. Development and execution can be done entirely using Docker.

## Requirements
- Docker
- Docker Compose

## Usage

Build and start the application with MySQL:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080/api`.

### Endpoints
- `POST /api/todos` - create a new todo (JSON `{ "title": "task" }`)
- `GET /api/todos` - list all todos
- `PUT /api/todos/{id}` - mark a todo as completed

## Project Structure

- `cmd/server` - application entry point
- `internal/entity` - domain models
- `internal/repository` - repository interfaces
- `internal/infrastructure/mysql` - MySQL implementation
- `internal/usecase` - business logic
- `internal/interface/controller` - HTTP handlers
- `internal/router` - Echo router setup

