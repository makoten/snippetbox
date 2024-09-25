package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	fmt.Fprint(w, "Hello from Snippetbox")
}

func createSnippetForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Show the create new snippet form...")
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// has to precede Write, otherwise Write sets 200 for statusCode
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, "Save a new snippet...")
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func main() {
	// I don't want to extend the default mux for security reasons
	mux := http.NewServeMux()

	// endpoints
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", viewSnippet)
	mux.HandleFunc("GET /snippet/create", createSnippetForm)
	mux.HandleFunc("POST /snippet/create", createSnippet)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
