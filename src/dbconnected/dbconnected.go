package main

import (
	"database/sql"
	"fmt"
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

	// query a single user from the db
	type quoteModel struct {
		id int 
		quote string
	}
	var q quoteModel

	singleQuoteQuery := "SELECT id, quote FROM quotes WHERE id = ?"
	
	// execute query by searching for the fields defined in struct
	if err := db.QueryRow(singleQuoteQuery, 1).Scan(&q.id, &q.quote); err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.quote)
}