package main

import (
	"errors"
	"fmt"
)

type TodoRepository struct {
	todos []*Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{todos: []*Todo{}}
}

func (r *TodoRepository) create(m *Todo) (*Todo, error) {
	r.todos = append(r.todos, m)
	return m, nil
}

func (r *TodoRepository) read(id string) (*Todo, error) {
	for _, element := range r.todos {
		if element.ID != id {
			continue
		}
		return element, nil
	}
	return &Todo{}, errors.New(fmt.Sprintf("UNKNOWN Todo: %s", id))
}

func (r *TodoRepository) update(m *Todo) {
	for index, element := range r.todos {
		if element.ID != m.ID {
			continue
		}
		r.todos[index] = m
	}
}

func (r *TodoRepository) delete(id string) (*Todo, error) {
	for index, element := range r.todos {
		if element.ID != id {
			continue
		}
		r.todos = append(r.todos[:index-1], r.todos[index+1:]...)
		return element, nil
	}
	return &Todo{}, errors.New(fmt.Sprintf("UNKNOWN Todo: %s", id))
}
