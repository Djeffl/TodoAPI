package main

import (
	"TodoAPI/database"
	"TodoAPI/routing"
	"log"
	"net/http"
)

func main() {
	//Invoke routing
	routing.InvokeRouter()

	//Connect to db
	database.OpenConnection()

	//Start server on port :8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
