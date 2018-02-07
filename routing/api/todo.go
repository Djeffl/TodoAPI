package api

import (
	"TodoAPI/app/todo"
	"TodoAPI/routing/models"
)

var todoRoutes = models.Routes{
	models.Route{
		Name:        "TodosIndex",
		Method:      "GET",
		Pattern:     "/api/todos",
		HandlerFunc: todo.GetTodos,
	},
	models.Route{
		Name:        "TodoDetail",
		Method:      "GET",
		Pattern:     "/api/todos/{todoId}",
		HandlerFunc: todo.GetTodo,
	},
	models.Route{
		Name:        "TodoCreate",
		Method:      "POST",
		Pattern:     "/api/todos",
		HandlerFunc: todo.CreateTodo,
	},
}
