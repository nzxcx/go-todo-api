package main

import (
	"go-todo-api/internal/delivery/http"
	"go-todo-api/internal/repository"
	"go-todo-api/internal/usecase"

	_ "go-todo-api/docs" // This is where the generated docs will be

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Todo API
// @version         1.0
// @description     A simple todo API using Go and Clean Architecture
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

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

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
