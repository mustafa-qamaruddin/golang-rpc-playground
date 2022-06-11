package server

import (
	"errors"
	"fmt"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	AssignedBy  string `json:"assignedBy"`
}

type ModelService interface {
	create(*interface{}) (*interface{}, error)
	read(*interface{}) (*interface{}, error)
	update(*interface{}) (*interface{}, error)
	delete(*interface{}) (*interface{}, error)
}

type TodoService struct {
	todos []*Todo
}

func NewTodoService() *TodoService {
	return &TodoService{todos: []*Todo{}}
}

func (s *TodoService) create(m *Todo) (*Todo, error) {
	s.todos = append(s.todos, m)
	return m, nil
}

func (s *TodoService) read(id string) (*Todo, error) {
	for _, element := range s.todos {
		if element.ID != id {
			continue
		}
		return element, nil
	}
	return &Todo{}, errors.New(fmt.Sprintf("UNKNOWN Todo: %s", id))
}

func (s *TodoService) update(m *Todo) {
	for index, element := range s.todos {
		if element.ID != m.ID {
			continue
		}
		s.todos[index] = m
	}
}

func (s *TodoService) delete(id string) (*Todo, error) {
	for index, element := range s.todos {
		if element.ID != id {
			continue
		}
		s.todos = append(s.todos[:index-1], s.todos[index+1:]...)
		return element, nil
	}
	return &Todo{}, errors.New(fmt.Sprintf("UNKNOWN Todo: %s", id))
}
