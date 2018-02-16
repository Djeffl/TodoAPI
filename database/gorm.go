package database

import (
	"TodoAPI/config"
	"TodoAPI/models"
	"fmt"
	"log"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func OpenConnection() {

	// Create connection string
	connString := fmt.Sprintf("server=%s;Database=%s;user id=%s;password=%s;port=%d",
		config.Server, config.Database, config.User, config.Password, config.Port)

	log.Println(connString)

	var err error
	db, err = gorm.Open("mssql", connString)
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected!")

	defer db.Close()

	initGorm()
}

func initGorm() {
	//Add all models
	db.AutoMigrate(&models.Todo{})
}
