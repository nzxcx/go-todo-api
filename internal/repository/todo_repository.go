package repository

import (
	"go-todo-api/internal/domain"
	"go-todo-api/internal/repository/models"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	db.AutoMigrate(&models.Todo{})
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	dbTodo := models.FromDomain(todo)
	if err := r.db.Create(dbTodo).Error; err != nil {
		return err
	}
	*todo = *dbTodo.ToDomain()
	return nil
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	var dbTodo models.Todo
	err := r.db.First(&dbTodo, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return dbTodo.ToDomain(), nil
}

func (r *todoRepository) GetAll() ([]*domain.Todo, error) {
	var dbTodos []models.Todo
	if err := r.db.Find(&dbTodos).Error; err != nil {
		return nil, err
	}

	todos := make([]*domain.Todo, len(dbTodos))
	for i, dbTodo := range dbTodos {
		todos[i] = dbTodo.ToDomain()
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	dbTodo := models.FromDomain(todo)
	result := r.db.Save(dbTodo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	*todo = *dbTodo.ToDomain()
	return nil
}

func (r *todoRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
} 
