package domain

import "time"

type Todo struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	GetByID(id uint) (*Todo, error)
	GetAll() ([]*Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}

type TodoUsecase interface {
	Create(todo *Todo) error
	GetByID(id uint) (*Todo, error)
	GetAll() ([]*Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
} 
