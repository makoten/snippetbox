package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a snippet..."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing a specific snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
