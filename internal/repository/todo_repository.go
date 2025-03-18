package repository

import (
	"errors"
	"sync"

	"go-todo-api/internal/domain"
)

type todoRepository struct {
	todos map[uint]*domain.Todo
	mutex sync.RWMutex
	nextID uint
}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{
		todos: make(map[uint]*domain.Todo),
		nextID: 1,
	}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	todo.ID = r.nextID
	r.todos[todo.ID] = todo
	r.nextID++
	return nil
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if todo, exists := r.todos[id]; exists {
		return todo, nil
	}
	return nil, errors.New("todo not found")
}

func (r *todoRepository) GetAll() ([]*domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos := make([]*domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}
	r.todos[todo.ID] = todo
	return nil
}

func (r *todoRepository) Delete(id uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.todos[id]; !exists {
		return errors.New("todo not found")
	}
	delete(r.todos, id)
	return nil
} 
