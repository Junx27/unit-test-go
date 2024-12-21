package todo

import (
	"testing"
)

func TestTodoManager(t *testing.T) {
	manager := NewTodoManager()

	manager.AddTodo("Task 1")
	manager.AddTodo("Task 2")
	if len(manager.GetTodos()) != 2 {
		t.Errorf("expected 2 todos, got %d", len(manager.GetTodos()))
	}

	err := manager.MarkDone(0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !manager.GetTodos()[0].Done {
		t.Errorf("expected todo to be marked as done")
	}

	err = manager.DeleteTodo(0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(manager.GetTodos()) != 1 {
		t.Errorf("expected 1 todo after deletion, got %d", len(manager.GetTodos()))
	}

	err = manager.MarkDone(10)
	if err == nil {
		t.Errorf("expected error for invalid index")
	}
	err = manager.DeleteTodo(10)
	if err == nil {
		t.Errorf("expected error for invalid index")
	}
}
