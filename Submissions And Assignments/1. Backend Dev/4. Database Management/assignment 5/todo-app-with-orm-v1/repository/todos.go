package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	return nil // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	return []model.Todo{}, nil // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	return nil // TODO: replace this
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	return nil // TODO: replace this
}
