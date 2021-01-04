package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)

func setQuote(w http.ResponseWriter, r *http.Request) {
	quote := "You and me, we're meant to be, roaming free, in harmony. One fine day, we'll fly away. Don't you know that Rome wasn't built in a day."
	result, err := db.Exec(`INSERT INTO quotes (quote) VALUES ?`, quote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]bool{"posted": false})
	}
	idResult = result.LastInsertID()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"result": idResult})
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	// query a single user from the db
	type quoteModel struct {
		id int 
		quote string
	}
	var q quoteModel

	singleQuoteQuery := "SELECT id, quote FROM quotes WHERE id = ?"
	
	// execute query by searching for the fields defined in struct
	if err := db.QueryRow(singleQuoteQuery, 1).Scan(&q.id, &q.quote); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"quote": "unable to retrieve"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"quote": q.quote})
}

func main() {
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/get-quote", getQuote).Methods("GET")
	r.HandleFunc("/api/v1/set-quote", setQuote).Methods("POST")
	r.HandleFunc("/api/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": false})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": true})
		}
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}