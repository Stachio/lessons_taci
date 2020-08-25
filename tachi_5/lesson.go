package main

import (
	"log"
	"net/http"
)

func tacisWebPronz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Taci's pronz website"))
}

func main() {
	/*
		http.Handle("/foo", fooHandler)

		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	http.HandleFunc("/pronz", tacisWebPronz)
	log.Fatal(http.ListenAndServe(":80", nil))
}

//() {} []
