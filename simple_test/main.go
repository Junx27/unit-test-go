package main

import (
	"fmt"
	"os"
	"unit-test-go/todo"
)

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic occurred:", r)
	}
}

func riskyOperation() {
	defer HandlePanic()
	fmt.Println("Starting risky operation...")
	panic("A serious error occurred!")
}

func main() {

	fmt.Println("Program started.")
	riskyOperation()
	fmt.Println("Program continues running after recovery.")

	f, _ := os.Open("non-existing-file")
	f.Close()

	const todoFormat = "- %s (done: %v)\n"
	manager := todo.NewTodoManager()

	manager.AddTodo("Learn Go")
	manager.AddTodo("Build a project")

	fmt.Println("Todos:")
	for _, t := range manager.GetTodos() {
		fmt.Printf(todoFormat, t.Description, t.Done)
	}
	err := manager.MarkDone(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nUpdated Todos:")
	for _, t := range manager.GetTodos() {
		fmt.Printf(todoFormat, t.Description, t.Done)
	}
	manager.DeleteTodo(0)
	fmt.Println("\nTodos after deletion:")
	for _, t := range manager.GetTodos() {
		fmt.Printf(todoFormat, t.Description, t.Done)
	}
}
