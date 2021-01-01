package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// create connection to the MySQL database
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// initalise the first connection to the database by pinging it and checking for error
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	quoteLengthQuery := `SELECT COUNT(id) FROM quotes`
	if _, err := db.Exec(quoteLengthQuery); err != nil {
		log.Fatal(err)
	}
}