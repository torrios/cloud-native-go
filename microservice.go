package main

import (
	"fmt"
	"json_marshalling/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go (Update)")
}