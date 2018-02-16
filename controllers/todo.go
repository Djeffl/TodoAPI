package controllers

import (
	"TodoAPI/database"
	"TodoAPI/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todos")

	// todos := []models.Todo{
	// 	{
	// 		ID:           "0",
	// 		Name:         "Taak1",
	// 		CreationDate: time.Now().Local(),
	// 		Completed:    false,
	// 	},
	// }
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var todo models.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(fmt.Errorf("Error: %v", err))
		return
	}
	log.Println("Todo: ")
	log.Println("Name: %s", todo.Name)
	log.Println("HasTime: %s", todo.HasTime)
	log.Println("HoursLeft: %s", todo.HoursLeft)

	database.CreateTodo(todo)

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
