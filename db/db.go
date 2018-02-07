package db

import (
	"TodoAPI/config"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func ConnectToDatabase() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		config.Server, config.User, config.Password, config.Port, config.Database)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// Create employee
	if err != nil {
		fmt.Printf("UUIDv4: %s\n", err)
	}

	createId, err := CreateTodo("Jake")
	fmt.Printf("Inserted ID: %d successfully.\n", createId)

	// // Read employees
	count, err := ReadEmployees()
	fmt.Printf("Read %d rows successfully.\n", count)
	//
	// // Update from database
	updateId, err := UpdateEmployee("Daniel", 1)
	fmt.Printf("Updated row with ID: %d successfully.\n", updateId)
	//
	// // Delete from database
	rows, err := DeleteEmployee("Daniel")
	fmt.Printf("Deleted %d rows successfully.\n", rows)

}

// parameters
//	query: Insert statement eg.: INSERT INTO [dbo].[Todos] (Name) VALUES (@aName);
//	args: Values eg.: { sql.Named("aName", aValueOfName) }
func Create(query string, args interface{}) (int64, error) {
	ctx := context.Background()

	var err error

	if db == nil {
		log.Fatal("No database connection established.")
	}

	// Check if database is alive
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	result, err := db.ExecContext(
		ctx,
		query,
		args,
	)

	if err != nil {
		log.Fatal("Error inserting new row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

func CreateTodo(name string) (int64, error) {
	// ctx := context.Background()
	// var err error
	//
	// if db == nil {
	// 	log.Fatal("What?")
	// }
	//
	// //Check if database is alive.
	// err = db.PingContext(ctx)
	// if err != nil {
	// 	log.Fatal("Error pinging database: " + err.Error())
	// }
	//
	// tsql := fmt.Sprintf("INSERT INTO [dbo].[Todos] (Name) VALUES (@Name);")
	//
	// //Execute non-query with named parameters
	// result, err := db.ExecContext(
	// 	ctx,
	// 	tsql,
	// 	sql.Named("Name", name))
	//
	// if err != nil {
	// 	log.Fatal("Error inserting new row: " + err.Error())
	// 	return -1, err
	// }
	//
	// return result.LastInsertId()

	query := "INSERT INTO [dbo].[Todos] (Name) VALUES (@Name);"

	return Create(query, sql.Named("Name", name))
}

func Read(query string) (*sql.Rows, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	// Execute query
	// return *sql.Rows, Error
	return db.QueryContext(ctx, query)
}

func Update(name string, id int) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("UPDATE [dbo].[Todos] SET Name = @Name WHERE Id = @Id")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", name),
		sql.Named("Id", id))
	if err != nil {
		log.Fatal("Error updating row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

func ReadEmployees() (int, error) {
	var count int = 0

	rows, err := Read("SELECT * FROM [dbo].[Todos]")

	if err != nil {
		log.Fatal("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	// Iterate through the result set.
	for rows.Next() {
		var name string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s\n", id, name)
		count++
	}
	return count, nil
}

// Update an employee's information
func UpdateEmployee(name string, id int) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("UPDATE [dbo].[Todos] SET Name = @Name WHERE Id = @Id")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", name),
		sql.Named("Id", id))
	if err != nil {
		log.Fatal("Error updating row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

// Delete an employee from database
func DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("DELETE FROM [dbo].[Todos] WHERE Name=@Name;")

	// Execute non-query with named parameters
	result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

// Gets and prints SQL Server version
func SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
