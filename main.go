package main

import (
	"fmt"
)

type todo struct {
	id int 
	title string 
	isCompleted bool
}

type todos []todo 

func main() {
	allTodos := readTodos()

	// Add a new todo
	allTodos.addNewTodo(todo{id: 3, title: "Wash Dishes", isCompleted: true})

	// Find a todo by id 
	foundedTodo := allTodos.readTodo(1)

	fmt.Println("Founded Todo: ", foundedTodo)

	// Remove specific todo by id
	// allTodos.removeTodo(1)

	// Change isCompleted flag of specific todo by id
	allTodos.changeStatus(2, false)

	// Print all the todos 
	allTodos.print()
}

func (t *todos) addNewTodo(nt todo) {
	*t = append(*t, nt)
}

func readTodos() todos {
	firstTodo := todo{
		id: 1,
		title: "Wash Dishes",
		isCompleted: false,
	}

	secondTodo := todo{
		id: 2,
		title: "Make Tea",
		isCompleted: true,
	}

	return todos{firstTodo, secondTodo}
}

func (t todos) readTodo(id int) todo {
	var foundedTodo todo

	for _, v := range t {
		if v.id == id {			
			foundedTodo = v
		} 
	}

	return todo(foundedTodo)
}

func (t todos) print() {
	for _, v := range t {
		fmt.Println(v)
	}
}

func (t *todos) removeTodo(id int) {
	for i, v := range (*t) {
		if v.id == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
		}
	}
}

func (t *todos) changeStatus(id int, status bool) {
	for i, v := range (*t) {
		if v.id == id {
			(*t)[i].isCompleted = status
		}
	}
}
