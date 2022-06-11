package main

type TodoService struct {
	repo *TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{repo: NewTodoRepository()}
}

func (s *TodoService) Create(in *Todo, out *Todo) error {
	out, err := s.repo.create(in)
	if err != nil {
		return err
	}
	return nil
}
