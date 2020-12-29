package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ping at %s\n", r.URL.Path)
	})

	fileServe := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServe))

	http.ListenAndServe(":8080", nil)
}