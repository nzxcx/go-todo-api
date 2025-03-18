package repository

import (
	"database/sql"
	"go-todo-api/internal/domain"
	"time"

	_ "github.com/lib/pq"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) domain.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	query := `
		INSERT INTO todos (title, description, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)
		RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		now,
	).Scan(&todo.ID)
}

func (r *todoRepository) GetByID(id uint) (*domain.Todo, error) {
	todo := &domain.Todo{}
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM todos
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return todo, err
}

func (r *todoRepository) GetAll() ([]*domain.Todo, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM todos`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*domain.Todo
	for rows.Next() {
		todo := &domain.Todo{}
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	query := `
		UPDATE todos
		SET title = $1, description = $2, completed = $3, updated_at = $4
		WHERE id = $5`

	result, err := r.db.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		time.Now(),
		todo.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *todoRepository) Delete(id uint) error {
	query := `DELETE FROM todos WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
} 
