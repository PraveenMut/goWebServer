package main

import (
	"database/sql"
	"go-sql-driver/mysql"
)

db, err := sql.Open("mysql", "mysql+pymysql://root:root@data:3306/mydatabase")