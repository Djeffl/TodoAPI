package api

import (
	"TodoAPI/controllers"
	"TodoAPI/routing/models"
)

var todoRoutes = models.Routes{
	models.Route{
		Name:        "TodosIndex",
		Method:      "GET",
		Pattern:     "/api/todos",
		HandlerFunc: controllers.GetTodos,
	},
	models.Route{
		Name:        "TodoDetail",
		Method:      "GET",
		Pattern:     "/api/todos/{todoId}",
		HandlerFunc: controllers.GetTodo,
	},
	models.Route{
		Name:        "TodoCreate",
		Method:      "POST",
		Pattern:     "/api/todos",
		HandlerFunc: controllers.CreateTodo,
	},
}
