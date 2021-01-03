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
		log.Fatal(err)
	}

	fmt.Println(q.quote)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

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
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": false})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"healthy": true})
		}
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}