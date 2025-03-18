#!/bin/sh

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
while ! nc -z postgres 5432; do
  sleep 0.1
done
echo "PostgreSQL is ready!"

# Run migrations
echo "Running migrations..."
migrate -path /app/migrations -database "${DATABASE_URL}" up

# Start the application
echo "Starting the application..."
exec "$@" 
