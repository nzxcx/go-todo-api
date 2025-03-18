package http

import (
	"net/http"
	"strconv"

	"go-todo-api/internal/domain"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, usecase domain.TodoUsecase) {
	handler := &TodoHandler{
		todoUsecase: usecase,
	}

	e.POST("/todos", handler.Create)
	e.GET("/todos", handler.GetAll)
	e.GET("/todos/:id", handler.GetByID)
	e.PUT("/todos/:id", handler.Update)
	e.DELETE("/todos/:id", handler.Delete)
}

func (h *TodoHandler) Create(c echo.Context) error {
	todo := new(domain.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.todoUsecase.Create(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAll(c echo.Context) error {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	todo, err := h.todoUsecase.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Todo not found",
		})
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	todo := new(domain.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	todo.ID = uint(id)
	if err := h.todoUsecase.Update(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	if err := h.todoUsecase.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
} 
