package main

import (
	"TodoAPI/db"
	"TodoAPI/routing"
	"log"
	"net/http"
)

// type Todo struct {
// 	ID           string    `json:"id,omitempty"`
// 	Name         string    `json:"name,omitempty"`
// 	CreationDate time.Time `json:"creationDate,omitempty"`
// 	FinishTime   time.Time `json:"finishDate"`
// 	Completed    bool      `json:"completed"`
// }
//
// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }

// type Routes []Route
//
// func NewRouter() *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		var handler http.Handler
// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)
//
// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(handler)
//
// 	}
// 	return router
// }

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
//
// 		inner.ServeHTTP(w, r)
//
// 		log.Printf(
// 			"%s\t%s\t%s\t%s",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			time.Since(start),
// 		)
// 	})
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome!")
// }
//
// func GetTodos(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Todos")
//
// 	todos := []Todo{
// 		{
// 			ID:           "0",
// 			Name:         "Taak1",
// 			CreationDate: time.Now().Local(),
// 			Completed:    false,
// 		},
// 	}
//
// 	json.NewEncoder(w).Encode(todos)
// 	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	//w.WriteHeader(http.StatusOK)
// 	// if err := json.NewEncoder(w).Encode(todos); err != nil {
// 	// 	panic(err)
// 	// }
// }
//
// func GetTodo(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoId := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show:", todoId)
// }
//
// func CreateTodo(w http.ResponseWriter, r *http.Request) {
// 	//var todo Todo
// 	//body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// if err := r.Body.Close(); err != nil {
// 	// 	panic(err)
// 	// }
// 	// if err := json.Unmarshal(body, &todo); err != nil {
// 	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	// 	w.WriteHeader(422) // unprocessable entity
// 	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	//}
//
// 	// t := RepoCreateTodo(todo)
// 	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	// w.WriteHeader(http.StatusCreated)
// 	// if err := json.NewEncoder(w).Encode(t); err != nil {
// 	// 	panic(err)
// 	// }
// }

// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/",
// 		Index,
// 	},
// 	Route{
// 		"TodosIndex",
// 		"GET",
// 		"/todos",
// 		GetTodos,
// 	},
// 	Route{
// 		"TodoDetail",
// 		"GET",
// 		"/todos/{todoId}",
// 		GetTodo,
// 	},
// 	Route{
// 		"TodoCreate",
// 		"POST",
// 		"/todos",
// 		CreateTodo,
// 	},
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello world!")
// }
//

func main() {
	//Invoke routing
	routing.InvokeRouter()

	//Connect to db
	db.ConnectToDatabase()

	//Start server on port :8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
