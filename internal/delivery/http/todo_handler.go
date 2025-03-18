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

// NewTodoHandler initializes the todo handler
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

// Create godoc
// @Summary      Create a new todo
// @Description  Create a new todo with the provided information
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      domain.Todo  true  "Todo object"
// @Success      201   {object}  domain.Todo
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /todos [post]
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

// GetAll godoc
// @Summary      List all todos
// @Description  Get all todos from the system
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.Todo
// @Failure      500  {object}  map[string]string
// @Router       /todos [get]
func (h *TodoHandler) GetAll(c echo.Context) error {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todos)
}

// GetByID godoc
// @Summary      Get a todo by ID
// @Description  Get a single todo by its ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  domain.Todo
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /todos/{id} [get]
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

// Update godoc
// @Summary      Update a todo
// @Description  Update a todo with the provided information
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Todo ID"
// @Param        todo  body      domain.Todo true  "Todo object"
// @Success      200   {object}  domain.Todo
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /todos/{id} [put]
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

// Delete godoc
// @Summary      Delete a todo
// @Description  Delete a todo by its ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /todos/{id} [delete]
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
