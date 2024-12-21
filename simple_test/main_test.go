package main

import (
	"os"
	"testing"
	"unit-test-go/todo"

	"github.com/stretchr/testify/assert"
)

func TestHandlePanic(t *testing.T) {

	defer HandlePanic()
	riskyOperation()

	t.Log("Panic telah berhasil di-handle.")
}

func TestTodoManager(t *testing.T) {

	manager := todo.NewTodoManager()

	manager.AddTodo("Learn Go")
	manager.AddTodo("Build a project")

	todos := manager.GetTodos()
	assert.Equal(t, 2, len(todos), "Jumlah todo harusnya 2")

	err := manager.MarkDone(1)
	assert.NoError(t, err, "Tidak seharusnya ada error ketika menandai todo sebagai selesai")

	assert.True(t, todos[1].Done, "Todo dengan ID 1 seharusnya sudah selesai")

	err = manager.MarkDone(100)
	assert.Error(t, err, "Seharusnya ada error ketika ID tidak ditemukan")

	manager.DeleteTodo(0)
	todos = manager.GetTodos()
	assert.Equal(t, 1, len(todos), "Jumlah todo harusnya 1 setelah penghapusan")
	assert.Equal(t, "Build a project", todos[0].Description, "Todo yang tersisa harusnya 'Build a project'")

	manager.DeleteTodo(100)
	todos = manager.GetTodos()
	assert.Equal(t, 1, len(todos), "Jumlah todo setelah delete invalid ID seharusnya tetap 1")
}

func TestFileOperation(t *testing.T) {

	_, err := os.Open("non-existing-file")
	assert.Error(t, err, "Seharusnya error ketika file tidak ditemukan")
}
