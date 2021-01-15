package services

import "github.com/onelzyugy/projects/learn-http/types"

var todos []types.Todo

// AddTodoItem is to add an items
func AddTodoItem() {
	firstItem := types.Todo{
		ItemID:      1,
		ItemName:    "my first item name",
		Description: "take out the trash",
	}
	secondItem := types.Todo{
		ItemID:      2,
		ItemName:    "my second item name",
		Description: "clean the house",
	}
	todos = append(todos, firstItem)
	todos = append(todos, secondItem)
}

// GetTodoItems retrieve all todos
func GetTodoItems() []types.Todo {
	return todos
}
