package main

import (
	"go-todo-api/internal/delivery/http"
	"go-todo-api/internal/repository"
	"go-todo-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize dependencies
	todoRepo := repository.NewTodoRepository()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	// Initialize handlers
	http.NewTodoHandler(e, todoUsecase)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
