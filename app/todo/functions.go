package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todos")

	todos := []Todo{
		{
			ID:           "0",
			Name:         "Taak1",
			CreationDate: time.Now().Local(),
			Completed:    false,
		},
	}

	json.NewEncoder(w).Encode(todos)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//var todo Todo
	//body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &todo); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	//}

	// t := RepoCreateTodo(todo)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}
