package main

import (
	"log"
	"encoding/json"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)

type DatabaseConnection struct {
	db *sql.DB
}

func initaliseDB() *sql.DB {
	databaseConnection, err := sql.Open("mysql", "root:root@(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	return databaseConnection
}

func (dbConnection *DatabaseConnection) setQuote(w http.ResponseWriter, r *http.Request) {
	quote := "You and me, we're meant to be, roaming free, in harmony. One fine day, we'll fly away. Don't you know that Rome wasn't built in a day."
	result, err := dbConnection.db.Exec(`INSERT INTO quotes (quote) VALUES ?`, quote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]bool{"posted": false})
	}
	idResult, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"result": idResult})
}

func (dbConnection *DatabaseConnection) getQuote(w http.ResponseWriter, r *http.Request) {
	// query a single user from the db
	type quoteModel struct {
		id int 
		quote string
	}
	var q quoteModel

	singleQuoteQuery := "SELECT id, quote FROM quotes WHERE id = ?"
	
	// execute query by searching for the fields defined in struct
	if err := dbConnection.db.QueryRow(singleQuoteQuery, 1).Scan(&q.id, &q.quote); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"quote": "unable to retrieve"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"quote": q.quote})
}

func main() {
	dbConnection := &DatabaseConnection{db: initaliseDB()}
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/get-quote", dbConnection.getQuote).Methods("GET")
	r.HandleFunc("/api/v1/set-quote", dbConnection.setQuote).Methods("POST")
	r.HandleFunc("/api/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := dbConnection.db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": false})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": true})
		}
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}