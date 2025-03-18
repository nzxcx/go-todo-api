# Go Todo API

A RESTful Todo API built with Go, using Clean Architecture, Echo framework, and PostgreSQL.

## Prerequisites

- Docker Desktop installed and running
- Git

## Quick Start

1. Clone the repository:

```bash
git clone git@github.com:nzxcx/go-todo-api.git
cd go-todo-api
```

2. Start the application:

```bash
docker compose up -d
```

The application will be available at:

- API: <http://localhost:8080>
- Swagger Documentation: <http://localhost:8080/swagger/index.html>
- pgAdmin: <http://localhost:5050>

## API Endpoints

- `POST /todos` - Create a new todo
- `GET /todos` - Get all todos
- `GET /todos/:id` - Get a specific todo
- `PUT /todos/:id` - Update a todo
- `DELETE /todos/:id` - Delete a todo

## Database Management with pgAdmin

1. Access pgAdmin at <http://localhost:5050>
2. Login credentials:

   - Email: <admin@admin.com>
   - Password: admin

3. Connect to the database:

   - Right-click on "Servers" → "Register" → "Server"
   - In the "General" tab:
     - Name: Todo DB (or any name you prefer)
   - In the "Connection" tab:
     - Host: postgres
     - Port: 5432
     - Database: tododb
     - Username: postgres
     - Password: postgres

4. View your tables:
   - Expand "Servers" → "Todo DB" → "Databases" → "tododb" → "Schemas" → "public" → "Tables"
   - You should see the `todos` table
   - Right-click on the `todos` table and select:
     - "View/Edit Data" → "All Rows" to see all todos
     - "Query Tool" to run custom SQL queries
     - "Properties" to see table structure

## Development

### Project Structure

```
.
├── internal/
│   ├── domain/      # Domain models and interfaces
│   ├── usecase/     # Business logic
│   ├── repository/  # Data access layer
│   └── delivery/    # HTTP handlers
├── migrations/      # Database migrations
├── scripts/         # Utility scripts
├── main.go         # Application entry point
└── docker-compose.yml
```

### Useful Commands

```bash
# Start the application
docker compose up -d

# View logs
docker compose logs -f

# Stop the application
docker compose down

# Rebuild and restart
docker compose up -d --build

# Check container status
docker compose ps
```

### Database Migrations

Migrations are automatically applied when the application starts. If you need to run migrations manually:

```bash
docker compose exec api migrate -path /app/migrations -database "${DATABASE_URL}" up
```

## API Testing

You can test the API using the Swagger UI at <http://localhost:8080/swagger/index.html> or using curl:

```bash
# Create a todo
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "description": "Study Go programming language"}'

# Get all todos
curl http://localhost:8080/todos

# Get a specific todo
curl http://localhost:8080/todos/1

# Update a todo
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "description": "Study Go programming language", "completed": true}'

# Delete a todo
curl -X DELETE http://localhost:8080/todos/1
```

## Troubleshooting

1. If containers fail to start:

   ```bash
   docker compose down -v  # Remove volumes
   docker compose up -d    # Start fresh
   ```

2. If you need to check database logs:

   ```bash
   docker compose logs postgres
   ```

3. If you need to access the database directly:

   ```bash
   docker compose exec postgres psql -U postgres -d tododb
   ```
