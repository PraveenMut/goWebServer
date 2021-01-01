package main

import (
	"database/sql"
	"log"
	"go-sql-driver/mysql"
)

db, err := sql.Open("mysql", "root:root@data:3306/mydatabase")

err := db.Ping()

if err != nil {
	log.Fatal(err)
}