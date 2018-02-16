package database

import (
	"TodoAPI/models"
	"log"
)

func CreateTodo(todo models.Todo) {
	val := db.Create(todo)
	log.Println("value create = : %s", val)
	//return db.First(&, 10)
}

func ReadAllTodos() {
	var todos []models.Todo
	db.First(&todos)
}

func UpdateTodo(todo models.Todo, fieldName string, value interface{}) {
	db.Model(&todo).Update(fieldName, value)
}

// func DeleteTodo(todo models.Todo) {
//   db.Delete(&todo)
// }

//  // Read
//  var product Product
//  db.First(&product, 1) // find product with id 1
//  db.First(&product, "code = ?", "L1212") // find product with code l1212
//
//  // Update - update product's price to 2000
//  db.Model(&product).Update("Price", 2000)
//
//  // Delete - delete product
//  db.Delete(&product)
