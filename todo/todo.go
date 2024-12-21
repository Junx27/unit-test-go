package todo

import "errors"

type Todo struct {
	Description string
	Done        bool
}

type TodoManager struct {
	todos []Todo
}

func NewTodoManager() *TodoManager {
	return &TodoManager{todos: []Todo{}}
}
func (m *TodoManager) AddTodo(description string) {
	m.todos = append(m.todos, Todo{Description: description, Done: false})
}

func (m *TodoManager) GetTodos() []Todo {
	return m.todos
}

func (m *TodoManager) MarkDone(index int) error {
	if index < 0 || index >= len(m.todos) {
		return errors.New("todo not found")
	}
	m.todos[index].Done = true
	return nil
}
func (m *TodoManager) DeleteTodo(index int) error {
	if index < 0 || index >= len(m.todos) {
		return errors.New("todo not found")
	}
	m.todos = append(m.todos[:index], m.todos[index+1:]...)
	return nil
}
